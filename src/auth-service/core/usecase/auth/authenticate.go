package authorize

import (
	"sync"

	authErrors "github.com/pedro-muniz/myPrice/auth/core/customerror/auth"
	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
	port "github.com/pedro-muniz/myPrice/auth/core/port/usecase/auth"
	publisher "github.com/pedro-muniz/myPrice/auth/core/port/publisher"
	repository "github.com/pedro-muniz/myPrice/auth/core/port/repository"
)

type Authenticate struct {
	AuthRepository repository.IAuthRepository
	AuthPublisher  publisher.AuthPublisher
}

var authenticateOnce sync.Once
var authenticateInstance *Authenticate

func GetAuthenticateInstance(authRepository repository.IAuthRepository,
	authPublisher publisher.AuthPublisher,
) *Authenticate {
	authenticateOnce.Do(func() {
		authenticateInstance = &Authenticate{AuthRepository: authRepository,
			AuthPublisher: authPublisher,
		}
	})

	return authenticateInstance
}

// Execute authenticates the client and returns an auth token
func (this *Authenticate) Execute(input *port.AuthenticateInput) (*port.AuthenticateOutput, error) {
	if input == nil {
		return nil, authErrors.InvalidAuthReference("")
	}

	// Convert input DTO to domain object
	auth := &domain.Auth{
		ClientId:     input.ClientId,
		ClientSecret: input.ClientSecret,
	}

	//Validate the token + secret on database
	authChan, repoErrChan := this.AuthRepository.Get(auth.ClientId, auth.ClientSecret)

	var client *domain.Auth
	var err error

	select {
	case client = <-authChan:
	case err = <-repoErrChan:
	}

	if err != nil {
		return nil, authErrors.ErrorGettingAuthDatabaseRecord(err.Error())
	}

	if client == nil {
		return nil, authErrors.ClientNotFound("")
	}

	authToken, err := auth.GetAuthToken()
	if err != nil {
		return nil, err
	}

	err = this.AuthPublisher.Publish(authToken.Token, authToken.ExpiringIn)
	if err != nil {
		return nil, err
	}

	// Convert domain object to output DTO
	return &port.AuthenticateOutput{
		ClientId:   authToken.ClientId,
		Token:      authToken.Token,
		ExpiringIn: authToken.ExpiringIn,
	}, nil
}

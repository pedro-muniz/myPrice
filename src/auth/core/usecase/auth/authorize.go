package authorize

import (
	"fmt"
	"sync"

	authErrors "github.com/pedro-muniz/myPrice/auth/core/customerror/auth"
	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
	publisher "github.com/pedro-muniz/myPrice/auth/core/port/publisher"
	repository "github.com/pedro-muniz/myPrice/auth/core/port/repository"
)

type Authorize struct {
	AuthRepository repository.AuthRepository
	AuthPublisher  publisher.AuthPublisher
}

var once sync.Once
var instance *Authorize

func GetInstance(authRepository repository.AuthRepository,
	authPublisher publisher.AuthPublisher,
) *Authorize {
	once.Do(func() {
		instance = &Authorize{AuthRepository: authRepository,
			AuthPublisher: authPublisher,
		}
	})

	return instance
}

//Authorize the client and return a the auth token
func (this *Authorize) Execute(auth *domain.Auth) (*domain.AuthToken, error) {
	//Validate the token + secret on database
	if auth == nil {
		return nil, authErrors.InvalidAuthReference("")
	}

	client, err := this.AuthRepository.Get(auth.ClientId, auth.ClientSecret)
	if err != nil {
		return nil, authErrors.ErrorGettingAuthDatabaseRecord(err.Error())
	}

	if client == nil {
		fmt.Println("TODO: implement client database")
		//return nil, authErrors.ClientNotFound("")
	}

	authToken, err := auth.GetAuthToken()
	if err != nil {
		return nil, err
	}

	err = this.AuthPublisher.Publish(authToken.Token, authToken.ExpiringAt)
	if err != nil {
		return nil, err
	}

	return authToken, nil
}

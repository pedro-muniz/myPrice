package authorize

import (
	authErrors "core/customerror/auth"
	domain "core/domain"
	publisher "core/port/publisher"
	repository "core/port/repository"
	"sync"
)

type Authorize struct {
	Auth           *domain.Auth
	AuthToken      *domain.AuthToken
	AuthRepository repository.AuthRepository
	AuthPublisher  publisher.AuthPublisher
}

var once sync.Once
var instance *Authorize

func GetInstance() *Authorize {
	once.Do(func() {
		instance = &Authorize{}
	})

	return instance
}

//Authorize the client and return a the auth token
func (this *Authorize) Execute() (*domain.AuthToken, error) {

	//Validate the token + secret on database
	if this.Auth == nil {
		return nil, authErrors.InvalidAuthReference("")
	}

	client, err := this.AuthRepository.Get(this.Auth.ClientId, this.Auth.ClientSecret)
	if err != nil {
		return nil, authErrors.ErrorGettingAuthDatabaseRecord(err.Error())
	}

	if client == nil {
		return nil, authErrors.ClientNotFound("")
	}

	authToken, err := this.Auth.GetAuthToken()
	if err != nil {
		return nil, err
	}

	err = this.AuthPublisher.Publish(authToken.Token, authToken.ExpiringAt)
	if err != nil {
		return nil, err
	}

	return authToken, nil
}

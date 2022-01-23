package usecases

import (
	authErrors "core/customErrors/auth"
	domain "core/domain"
	port "core/port/repository"
)

type Authorize struct {
	Auth           *domain.Auth
	AuthRepository port.AuthRepository
}

//Authorize the client and return a the auth token
func (this *Authorize) Execute() (string, error) {
	//Validate the token + secret on database
	if this.Auth == nil {
		return "", authErrors.InvalidAuthReference("")
	}

	client, err := this.AuthRepository.Get(this.Auth.ClientId, this.Auth.ClientSecret)

	if err != nil {
		return "", authErrors.ErrorGettingAuthDatabaseRecord(err.Error())
	}

	if client == nil {
		return "", authErrors.ClientNotFound("")
	}

	return this.Auth.GenerateToken()
}

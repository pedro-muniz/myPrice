package domain

import (
	authErrors "core/customerror/auth"
	"crypto/rand"
	"fmt"
	"time"
)

type Auth struct {
	ClientId     string
	ClientSecret string
	GrantType    string
}

type AuthToken struct {
	ClientId   string
	Token      string
	ExpiringAt time.Time
}

//Generate single user token
func (this *Auth) generateToken() (string, error) {
	if len(this.ClientId) <= 0 {
		return "", authErrors.InvalidClientId("")
	}

	token := make([]byte, 16)
	rand.Read(token)
	return fmt.Sprintf("%x", token), nil
}

//Populate token struct
func (this *Auth) GetAuthToken() (*AuthToken, error) {
	token, err := this.generateToken()
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		ClientId:   this.ClientId,
		Token:      token,
		ExpiringAt: time.Now(),
	}, nil
}

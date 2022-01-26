package domain

import (
	"crypto/rand"
	"fmt"
	"time"

	authErrors "github.com/pedro-muniz/myPrice/auth/core/customerror/auth"
)

type Auth struct {
	ClientId     string
	ClientSecret string
	GrantType    string
}

type AuthToken struct {
	ClientId   string
	Token      string
	ExpiringAt time.Duration
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

	expiringAt := 1 + time.Hour

	return &AuthToken{
		ClientId:   this.ClientId,
		Token:      token,
		ExpiringAt: expiringAt,
	}, nil
}

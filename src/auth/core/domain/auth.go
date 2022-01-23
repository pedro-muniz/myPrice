package domain

import (
	authErrors "core/customErrors/auth"
	"crypto/rand"
	"fmt"
)

type Auth struct {
	ClientId     string
	ClientSecret string
	GrantType    string
}

//Generate single user token
func (this *Auth) GenerateToken() (string, error) {
	if len(this.ClientId) <= 0 {
		return "", authErrors.InvalidClientId("")
	}

	token := make([]byte, 16)
	rand.Read(token)
	return fmt.Sprintf("%x", token), nil
}

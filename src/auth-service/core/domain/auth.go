package domain

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authErrors "github.com/pedro-muniz/myPrice/auth/core/customerror/auth"
)

type Auth struct {
	ClientId     string
	ClientName   string
	ClientEmail  string
	ClientSecret string
	GrantType    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastLoginAt  time.Time
}

type AuthToken struct {
	ClientId   string
	Token      string
	ExpiringIn time.Duration
}

// Generate single user token
func (this *Auth) generateToken(expiringIn time.Duration) (string, error) {
	if len(this.ClientId) <= 0 {
		return "", authErrors.InvalidClientId("")
	}

	claims := jwt.MapClaims{
		"client_id": this.ClientId,
		"exp":       time.Now().Add(expiringIn).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("auth_secret"))

	if this.ClientSecret != "" {
		secret = []byte(this.ClientSecret)
	}

	return token.SignedString(secret)
}

// Populate token struct
func (this *Auth) GetAuthToken() (*AuthToken, error) {
	expiringIn := 1 + time.Hour
	token, err := this.generateToken(expiringIn)
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		ClientId:   this.ClientId,
		Token:      token,
		ExpiringIn: expiringIn,
	}, nil
}

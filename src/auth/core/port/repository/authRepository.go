package repository

import (
	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
)

//Interface to get users from database
type AuthRepository interface {
	Get(clientId string, clientSecret string) (*domain.Auth, error)
	Save(auth *domain.Auth) error
}

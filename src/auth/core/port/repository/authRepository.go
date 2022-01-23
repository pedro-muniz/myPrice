package repository

import (
	domain "core/domain"
)

//Interface to get users from database
type AuthRepository interface {
	Get(clientId string, clientSecret string) (*domain.Auth, error)
	Save(auth *domain.Auth) error
}

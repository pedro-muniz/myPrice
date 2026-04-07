package repository

import (
	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
)

//Interface to get users from database
type IAuthRepository interface {
	Get(email string, password string) (<-chan *domain.Auth, <-chan error)
	Save(auth *domain.Auth) <-chan error
}

package auth

import (
	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
)

// IAuthorize defines the behavior for the authentication use case
type IAuthorize interface {
	Execute(auth *domain.Auth) (*domain.AuthToken, error)
	Validate(token string) (*domain.AuthToken, error)
}

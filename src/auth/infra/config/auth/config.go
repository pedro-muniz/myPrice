package auth

import (
	authRepo "github.com/pedro-muniz/myPrice/auth/infra/persistence/repository/neo4j/auth"
)

type AuthConfig struct {
	repo *authRepo.AuthRepository
}

func (this *AuthConfig) CreateAuthorizeUseCase() {
	this.repo = authRepo.GetInstance()
}

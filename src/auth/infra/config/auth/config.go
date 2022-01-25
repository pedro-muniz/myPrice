package auth

import (
	authRepository "infra/persistence/repository/neo4j/auth"
)

type AuthConfig struct {
	repo AuthRepository
}

func (this *AuthConfig) CreateAuthorizeUseCase() {
	this.repo = authRepository.GetInstance()
}

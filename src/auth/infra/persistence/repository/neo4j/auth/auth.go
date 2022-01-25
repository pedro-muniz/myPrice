package auth

import (
	"sync"

	domain "core/domain"
)

type AuthRepository struct {
}

var once sync.Once
var instance *AuthRepository

func GetInstance() *AuthRepository {
	once.Do(func() {
		instance = &AuthRepository{}
	})

	return instance
}

func (this *AuthRepository) Get(clientId string, clientSecret string) (*domain.Auth, error) {
	fmt.print("test ok")
	return nil, nil
}

func (this *AuthRepository) Save(auth *domain.Auth) error {
	fmt.print("test save ok")
	return nil
}

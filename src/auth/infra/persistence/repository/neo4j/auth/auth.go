package auth

import (
	"fmt"
	"sync"

	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
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
	fmt.Println("test ok")
	return nil, nil
}

func (this *AuthRepository) Save(auth *domain.Auth) error {
	fmt.Println("test save ok")
	return nil
}

package auth

import (
	"fmt"
	"sync"

	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
)

type AuthModel struct {
	Id       uint64 //database id
	Name     string //client name
	Email    string //email to login. I.e: email
	Password string //client password or secret
	Roles    string //client permissions
}

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

func (this *AuthRepository) Get(email string, password string) (*domain.Auth, error) {
	fmt.Println("test ok")
	//connect to database
	//look for auth node
	return nil, nil
}

func (this *AuthRepository) Save(auth *domain.Auth) error {
	//connect to database
	//save a auth node
	fmt.Println("test save ok")
	return nil
}

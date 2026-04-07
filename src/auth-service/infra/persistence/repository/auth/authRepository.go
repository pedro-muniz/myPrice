package auth

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
)

type AuthModel struct {
	Id          string
	Name        string
	Email       string
	Password    string
	Roles       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LastLoginAt time.Time
}

type IDAO interface {
	Read(query string, args ...interface{}) (*sql.Rows, error)
	Insert(authModel AuthModel) (sql.Result, error)
	Update(authModel AuthModel) (sql.Result, error)
	Delete(id string) (sql.Result, error)
}

type AuthRepository struct {
	dao IDAO
}

var once sync.Once
var instance *AuthRepository

func GetInstance(dao IDAO) *AuthRepository {
	once.Do(func() {
		instance = &AuthRepository{
			dao: dao,
		}
	})

	return instance
}

func (this *AuthRepository) Get(email string, password string) (<-chan *domain.Auth, <-chan error) {
	authChan := make(chan *domain.Auth, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(authChan)
		defer close(errChan)

		query := `
			SELECT auth_id as id, name, email, password, roles, created_at, updated_at, last_login_at
			FROM user_logins.auths 
			WHERE auth_id = $1 AND password = $2
		`
		rows, err := this.dao.Read(query, email, password)
		if err != nil {
			errChan <- fmt.Errorf("error reading from db: %w", err)
			return
		}
		defer rows.Close()

		if rows.Next() {
			var model AuthModel
			err := rows.Scan(&model.Id, &model.Name, &model.Email, &model.Password,
				&model.Roles, &model.CreatedAt, &model.UpdatedAt, &model.LastLoginAt)

			if err != nil {
				errChan <- fmt.Errorf("error scanning row: %w", err)
				return
			}

			authChan <- &domain.Auth{
				ClientName:   model.Name,
				ClientId:     model.Id,
				ClientEmail:  model.Email,
				ClientSecret: model.Password,
				GrantType:    model.Roles,
				CreatedAt:    model.CreatedAt,
				UpdatedAt:    model.UpdatedAt,
				LastLoginAt:  model.LastLoginAt,
			}
			return
		}

		errChan <- fmt.Errorf("user not found")
	}()

	return authChan, errChan
}

func (this *AuthRepository) Save(auth *domain.Auth) <-chan error {
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)

		fmt.Println("saving to postgres")

		model := AuthModel{
			Id:          auth.ClientId,
			Name:        auth.ClientName,
			Email:       auth.ClientEmail,
			Password:    auth.ClientSecret,
			Roles:       auth.GrantType,
			CreatedAt:   auth.CreatedAt,
			UpdatedAt:   auth.UpdatedAt,
			LastLoginAt: auth.LastLoginAt,
		}

		_, err := this.dao.Insert(model)
		if err != nil {
			errChan <- fmt.Errorf("error writing to db: %w", err)
		}
	}()

	return errChan
}

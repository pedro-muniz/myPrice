package authorize

import (
	"fmt"
	"os"
	"sync"

	"github.com/golang-jwt/jwt/v5"
	authErrors "github.com/pedro-muniz/myPrice/auth/core/customerror/auth"
	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
	publisher "github.com/pedro-muniz/myPrice/auth/core/port/publisher"
	repository "github.com/pedro-muniz/myPrice/auth/core/port/repository"
)

type Authorize struct {
	AuthRepository repository.IAuthRepository
	AuthPublisher  publisher.AuthPublisher
}

var once sync.Once
var instance *Authorize

func GetInstance(authRepository repository.IAuthRepository,
	authPublisher publisher.AuthPublisher,
) *Authorize {
	once.Do(func() {
		instance = &Authorize{AuthRepository: authRepository,
			AuthPublisher: authPublisher,
		}
	})

	return instance
}

// Authorize the client and return a the auth token
func (this *Authorize) Execute(auth *domain.Auth) (*domain.AuthToken, error) {
	//Validate the token + secret on database
	if auth == nil {
		return nil, authErrors.InvalidAuthReference("")
	}

	authChan, repoErrChan := this.AuthRepository.Get(auth.ClientId, auth.ClientSecret)

	var client *domain.Auth
	var err error

	select {
	case client = <-authChan:
	case err = <-repoErrChan:
	}

	if err != nil {
		return nil, authErrors.ErrorGettingAuthDatabaseRecord(err.Error())
	}

	if client == nil {
		return nil, authErrors.ClientNotFound("")
	}

	authToken, err := auth.GetAuthToken()
	if err != nil {
		return nil, err
	}

	err = this.AuthPublisher.Publish(authToken.Token, authToken.ExpiringIn)
	if err != nil {
		return nil, err
	}

	return authToken, nil
}

// Validate the jwt token and return the auth token
func (this *Authorize) Validate(token string) (*domain.AuthToken, error) {
	// 1. Get from publisher (Redis)
	authToken, err := this.AuthPublisher.Get(token)
	if err != nil {
		return nil, authErrors.InvalidAuthReference("Token not found or invalid")
	}

	if authToken == nil {
		return nil, authErrors.InvalidAuthReference("Token not found")
	}

	// 2. Parse and validate the JWT
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("auth_secret")), nil
	})

	if err != nil || !parsedToken.Valid {
		return nil, authErrors.InvalidAuthReference("Invalid or expired token")
	}

	// 3. Extract client_id from claims
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		if clientId, ok := claims["client_id"].(string); ok {
			authToken.ClientId = clientId
		}
	}

	return authToken, nil
}

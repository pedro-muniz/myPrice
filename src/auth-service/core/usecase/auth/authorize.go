package authorize

import (
	"fmt"
	"os"
	"sync"

	"github.com/golang-jwt/jwt/v5"
	authErrors "github.com/pedro-muniz/myPrice/auth/core/customerror/auth"
	port "github.com/pedro-muniz/myPrice/auth/core/port/usecase/auth"
	publisher "github.com/pedro-muniz/myPrice/auth/core/port/publisher"
)

type Authorize struct {
	AuthPublisher publisher.AuthPublisher
}

var once sync.Once
var instance *Authorize

func GetInstance(authPublisher publisher.AuthPublisher,
) *Authorize {
	once.Do(func() {
		instance = &Authorize{AuthPublisher: authPublisher}
	})

	return instance
}

// Execute validates the jwt token and returns the auth token
func (this *Authorize) Execute(token string) (*port.AuthorizeOutput, error) {
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

	// 3. Extract client_id from claims and build output DTO
	output := &port.AuthorizeOutput{
		Token:      authToken.Token,
		ExpiringIn: authToken.ExpiringIn,
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		if clientId, ok := claims["client_id"].(string); ok {
			output.ClientId = clientId
		}
	}

	return output, nil
}

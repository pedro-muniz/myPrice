package publisher

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
)

//Interface to get users from database
type AuthPublisher interface {
	Publish(token string, expiringIn time.Duration) error
	Get(token string) (*domain.AuthToken, error)
}

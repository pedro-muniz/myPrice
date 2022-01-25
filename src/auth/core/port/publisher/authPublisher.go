package publisher

import (
	"time"
)

//Interface to get users from database
type AuthPublisher interface {
	Publish(token string, expiringAt time.Time) error
}

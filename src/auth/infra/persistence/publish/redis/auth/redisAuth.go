package auth

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
)

type RedisAuth struct {
	Client *redis.Client
}

var once sync.Once
var instance *RedisAuth

func GetInstance(client *redis.Client) *RedisAuth {
	once.Do(func() {
		instance = &RedisAuth{Client: client}
	})

	return instance
}

var (
	ctx = context.Background()
)

//Set the token and expiringIn to redis
func (this *RedisAuth) Publish(token string, expiringIn time.Duration) error {
	// Last argument is expiration time.
	err := this.Client.Set(ctx, token, expiringIn.String(), expiringIn).Err()

	if err != nil {
		return err
	}

	return nil
}

//Get the token and expiringIn to redis
func (this *RedisAuth) Get(token string) (*domain.AuthToken, error) {
	val, err := this.Client.Get(ctx, token).Result()
	if err != nil {
		return nil, err
	}

	duration, _ := time.ParseDuration(val)
	return &domain.AuthToken{
		Token:      token,
		ExpiringIn: duration,
	}, nil
}

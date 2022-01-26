package auth

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	redismock "github.com/go-redis/redismock/v8"
)

func TestNewsInfoForCache(t *testing.T) {
	db, mock := redismock.NewClientMock()

	duration := 1 + time.Hour
	token := make([]byte, 16)
	rand.Read(token)

	key := fmt.Sprintf("%x", token)

	mock.ExpectSet(key, duration, duration).SetVal("OK")
	mock.ExpectGet(key).SetVal("value")

	redisAuth := &RedisAuth{Client: db}

	err := redisAuth.Publish(key, duration)
	if err != nil {
		t.Errorf("Error running redis-mock %s", err.Error())
	}
}

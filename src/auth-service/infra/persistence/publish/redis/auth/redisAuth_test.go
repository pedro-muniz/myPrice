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

	token := make([]byte, 16)
	rand.Read(token)
	key := fmt.Sprintf("%x", token)

	expiringIn := time.Hour + 1

	mock.ExpectSet(key, expiringIn.String(), expiringIn).SetVal("OK")

	redisAuth := &RedisAuth{Client: db}

	err := redisAuth.Publish(key, expiringIn)
	if err != nil {
		t.Errorf("Error running redis-mock %s", err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met: %s", err.Error())
	}
}

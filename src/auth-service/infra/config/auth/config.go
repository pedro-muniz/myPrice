package auth

import (
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	authUc "github.com/pedro-muniz/myPrice/auth/core/usecase/auth"
	r "github.com/pedro-muniz/myPrice/auth/infra/persistence/publish/redis/auth"
	authRepo "github.com/pedro-muniz/myPrice/auth/infra/persistence/repository/auth"
	postgres "github.com/pedro-muniz/myPrice/auth/infra/persistence/repository/postgres"
)

func createPostgresDao() *postgres.DAO {
	port := 5432
	if envPort := os.Getenv("postgres_port"); envPort != "" {
		if p, err := strconv.Atoi(envPort); err == nil {
			port = p
		}
	}

	return &postgres.DAO{
		Host:     os.Getenv("postgres_host"),
		Port:     port,
		User:     os.Getenv("postgres_user"),
		Password: os.Getenv("postgres_pass"),
		DbName:   os.Getenv("postgres_db"),
	}
}

func createRedisClient() *redis.Client {
	addr := os.Getenv("redis_addr")
	if addr == "" {
		addr = "localhost:6379"
	}

	db := 0
	if envDb := os.Getenv("redis_db"); envDb != "" {
		if d, err := strconv.Atoi(envDb); err == nil {
			db = d
		}
	}

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: os.Getenv("redis_pass"),
		DB:       db,
	})
}

var (
	repo      = authRepo.GetInstance(createPostgresDao())
	redisAuth = r.GetInstance(createRedisClient())
)

func CreateAuthorizeUseCase() *authUc.Authorize {
	authorizeUc := authUc.GetInstance(repo, redisAuth)
	return authorizeUc
}

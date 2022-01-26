package auth

import (
	"github.com/go-redis/redis/v8"
	authUc "github.com/pedro-muniz/myPrice/auth/core/usecase/auth"
	r "github.com/pedro-muniz/myPrice/auth/infra/persistence/publish/redis/auth"
	authRepo "github.com/pedro-muniz/myPrice/auth/infra/persistence/repository/neo4j/auth"
)

var (
	repo      = authRepo.GetInstance()
	redisAuth = r.GetInstance(redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}))
)

func CreateAuthorizeUseCase() *authUc.Authorize {
	authorizeUc := authUc.GetInstance(repo, redisAuth)
	return authorizeUc
}

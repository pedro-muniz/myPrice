package auth

import (
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pedro-muniz/myPrice/auth/core/port/publisher"
	"github.com/pedro-muniz/myPrice/auth/core/port/repository"
	useCasePort "github.com/pedro-muniz/myPrice/auth/core/port/usecase/auth"
	authUc "github.com/pedro-muniz/myPrice/auth/core/usecase/auth"
	"github.com/pedro-muniz/myPrice/auth/infra/delivery/controllers"
	r "github.com/pedro-muniz/myPrice/auth/infra/persistence/publish/redis/auth"
	authRepo "github.com/pedro-muniz/myPrice/auth/infra/persistence/repository/auth"
	postgres "github.com/pedro-muniz/myPrice/auth/infra/persistence/repository/postgres"
)

// ProviderSet for Wire dependency injection
var ProviderSet = wire.NewSet(
	ProvidePostgresDAO,
	ProvideRedisClient,
	ProvideAuthRepository,
	ProvideRedisAuth,
	ProvideAuthorizeUseCase,
	ProvideAuthController,
	// Interface bindings
	wire.Bind(new(authRepo.IDAO), new(*postgres.DAO)),
	wire.Bind(new(repository.IAuthRepository), new(*authRepo.AuthRepository)),
	wire.Bind(new(publisher.AuthPublisher), new(*r.RedisAuth)),
	wire.Bind(new(useCasePort.IAuthorize), new(*authUc.Authorize)),
)

func ProvidePostgresDAO() *postgres.DAO {
	port := 5432
	if envPort := os.Getenv("postgres_port"); envPort != "" {
		if p, err := strconv.Atoi(envPort); err == nil {
			port = p
		}
	}

	return postgres.GetInstance(
		os.Getenv("postgres_host"),
		port,
		os.Getenv("postgres_user"),
		os.Getenv("postgres_pass"),
		os.Getenv("postgres_db"),
	)
}

func ProvideRedisClient() *redis.Client {
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

func ProvideAuthRepository(dao authRepo.IDAO) *authRepo.AuthRepository {
	return authRepo.GetInstance(dao)
}

func ProvideRedisAuth(client *redis.Client) *r.RedisAuth {
	return r.GetInstance(client)
}

func ProvideAuthorizeUseCase(repo repository.IAuthRepository, redisAuth publisher.AuthPublisher) *authUc.Authorize {
	return authUc.GetInstance(repo, redisAuth)
}

func ProvideAuthController(useCase useCasePort.IAuthorize) *controllers.AuthController {
	return &controllers.AuthController{UseCase: useCase}
}

//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	authUc "github.com/pedro-muniz/myPrice/auth/core/usecase/auth"
	authConfig "github.com/pedro-muniz/myPrice/auth/infra/config/auth"
	"github.com/pedro-muniz/myPrice/auth/infra/delivery/controllers"
)

func InitializeAuthorizeUseCase() *authUc.Authorize {
	wire.Build(authConfig.ProviderSet)
	return nil
}

func InitializeAuthController() *controllers.AuthController {
	wire.Build(authConfig.ProviderSet)
	return nil
}

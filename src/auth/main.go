package main

import (
	"fmt"

	"github.com/pedro-muniz/myPrice/auth/core/domain"
	authConfig "github.com/pedro-muniz/myPrice/auth/infra/config/auth"
)

func main() {
	authorizeUc := authConfig.CreateAuthorizeUseCase()
	clientAuthentication := &domain.Auth{ClientId: "pedro", ClientSecret: "teste"}
	clientToken, err := authorizeUc.Execute(clientAuthentication)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", clientToken)

	clientAuthentication = &domain.Auth{ClientId: "test2", ClientSecret: "test3"}
	clientToken, err = authorizeUc.Execute(clientAuthentication)
	fmt.Printf("%+v\n", clientToken)
}

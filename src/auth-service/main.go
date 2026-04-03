package main

import (
	"fmt"

	"github.com/pedro-muniz/myPrice/auth/core/domain"
	authConfig "github.com/pedro-muniz/myPrice/auth/infra/config/auth"
)

func main() {
	authorizeUc := authConfig.CreateAuthorizeUseCase()
	clientAuthentication := &domain.Auth{ClientId: "726a28ba-b07b-4848-808d-6b345448357e", ClientSecret: "teste"}
	clientToken, err := authorizeUc.Execute(clientAuthentication)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", clientToken)

	clientAuthentication = &domain.Auth{ClientId: "2f546b9d-4143-4f00-accb-a2263ed52007", ClientSecret: "test3"}
	clientToken, err = authorizeUc.Execute(clientAuthentication)
	fmt.Printf("%+v\n", clientToken)
}

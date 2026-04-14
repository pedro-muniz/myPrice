package converters

import (
	port "github.com/pedro-muniz/myPrice/auth/core/port/usecase/auth"
	"github.com/pedro-muniz/myPrice/auth/infra/delivery/models"
)

func ToAuthenticateInput(request models.AuthRequest) *port.AuthenticateInput {
	return &port.AuthenticateInput{
		ClientId:     request.ClientId,
		ClientSecret: request.ClientSecret,
	}
}

func AuthenticateOutputToResponse(output *port.AuthenticateOutput) *models.AuthResponse {
	return &models.AuthResponse{
		Token:      output.Token,
		ExpiringIn: output.ExpiringIn.String(),
	}
}

func AuthorizeOutputToResponse(output *port.AuthorizeOutput) *models.AuthResponse {
	return &models.AuthResponse{
		Token:      output.Token,
		ExpiringIn: output.ExpiringIn.String(),
	}
}

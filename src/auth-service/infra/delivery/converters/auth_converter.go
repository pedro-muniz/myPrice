package converters

import (
	"github.com/pedro-muniz/myPrice/auth/core/domain"
	"github.com/pedro-muniz/myPrice/auth/infra/delivery/models"
)

func ToDomain(request models.AuthRequest) *domain.Auth {
	return &domain.Auth{
		ClientId:     request.ClientId,
		ClientSecret: request.ClientSecret,
	}
}

func ToResponse(token *domain.AuthToken) *models.AuthResponse {
	return &models.AuthResponse{
		Token:      token.Token,
		ExpiringIn: token.ExpiringIn.String(),
	}
}

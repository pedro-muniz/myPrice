package converters

import (
	"testing"
	"time"

	"github.com/pedro-muniz/myPrice/auth/core/domain"
	"github.com/pedro-muniz/myPrice/auth/infra/delivery/models"
)

func TestToDomain(t *testing.T) {
	request := models.AuthRequest{
		ClientId:     "test-client",
		ClientSecret: "test-secret",
	}

	result := ToDomain(request)

	if result.ClientId != request.ClientId {
		t.Errorf("Expected ClientId %s, got %s", request.ClientId, result.ClientId)
	}

	if result.ClientSecret != request.ClientSecret {
		t.Errorf("Expected ClientSecret %s, got %s", request.ClientSecret, result.ClientSecret)
	}
}

func TestToResponse(t *testing.T) {
	token := &domain.AuthToken{
		Token:      "test-token",
		ExpiringIn: time.Hour,
	}

	result := ToResponse(token)

	if result.Token != token.Token {
		t.Errorf("Expected Token %s, got %s", token.Token, result.Token)
	}

	expectedExp := "1h0m0s"
	if result.ExpiringIn != expectedExp {
		t.Errorf("Expected ExpiringIn %s, got %s", expectedExp, result.ExpiringIn)
	}
}

package converters

import (
	"testing"
	"time"

	port "github.com/pedro-muniz/myPrice/auth/core/port/usecase/auth"
	"github.com/pedro-muniz/myPrice/auth/infra/delivery/models"
)

func TestToAuthenticateInput(t *testing.T) {
	request := models.AuthRequest{
		ClientId:     "test-client",
		ClientSecret: "test-secret",
	}

	result := ToAuthenticateInput(request)

	if result.ClientId != request.ClientId {
		t.Errorf("Expected ClientId %s, got %s", request.ClientId, result.ClientId)
	}

	if result.ClientSecret != request.ClientSecret {
		t.Errorf("Expected ClientSecret %s, got %s", request.ClientSecret, result.ClientSecret)
	}
}

func TestAuthenticateOutputToResponse(t *testing.T) {
	output := &port.AuthenticateOutput{
		Token:      "test-token",
		ExpiringIn: time.Hour,
	}

	result := AuthenticateOutputToResponse(output)

	if result.Token != output.Token {
		t.Errorf("Expected Token %s, got %s", output.Token, result.Token)
	}

	expectedExp := "1h0m0s"
	if result.ExpiringIn != expectedExp {
		t.Errorf("Expected ExpiringIn %s, got %s", expectedExp, result.ExpiringIn)
	}
}

func TestAuthorizeOutputToResponse(t *testing.T) {
	output := &port.AuthorizeOutput{
		Token:      "test-token",
		ExpiringIn: time.Hour,
	}

	result := AuthorizeOutputToResponse(output)

	if result.Token != output.Token {
		t.Errorf("Expected Token %s, got %s", output.Token, result.Token)
	}

	expectedExp := "1h0m0s"
	if result.ExpiringIn != expectedExp {
		t.Errorf("Expected ExpiringIn %s, got %s", expectedExp, result.ExpiringIn)
	}
}

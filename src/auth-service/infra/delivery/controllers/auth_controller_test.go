package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
	"github.com/pedro-muniz/myPrice/auth/infra/delivery/models"
)

type MockAuthorize struct {
	FakeExecute  func(auth *domain.Auth) (*domain.AuthToken, error)
	FakeValidate func(token string) (*domain.AuthToken, error)
}

func (m *MockAuthorize) Execute(auth *domain.Auth) (*domain.AuthToken, error) {
	if m.FakeExecute != nil {
		return m.FakeExecute(auth)
	}
	return nil, nil
}

func (m *MockAuthorize) Validate(token string) (*domain.AuthToken, error) {
	if m.FakeValidate != nil {
		return m.FakeValidate(token)
	}
	return nil, nil
}

func TestAuthController_Authorize(t *testing.T) {
	t.Run("should return 200 and token on success", func(t *testing.T) {
		mockUseCase := &MockAuthorize{
			FakeExecute: func(auth *domain.Auth) (*domain.AuthToken, error) {
				return &domain.AuthToken{Token: "test-token", ExpiringIn: time.Hour}, nil
			},
		}
		controller := &AuthController{UseCase: mockUseCase}

		body, _ := json.Marshal(models.AuthRequest{ClientId: "id", ClientSecret: "secret"})
		req, _ := http.NewRequest(http.MethodPost, "/authorize", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		controller.Authorize(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", rr.Code)
		}

		var resp models.AuthResponse
		json.Unmarshal(rr.Body.Bytes(), &resp)
		if resp.Token != "test-token" {
			t.Errorf("Expected token test-token, got %s", resp.Token)
		}
	})

	t.Run("should return 405 for invalid method", func(t *testing.T) {
		controller := &AuthController{UseCase: &MockAuthorize{}}
		req, _ := http.NewRequest(http.MethodGet, "/authorize", nil)
		rr := httptest.NewRecorder()

		controller.Authorize(rr, req)

		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status 405, got %v", rr.Code)
		}
	})

	t.Run("should return 401 on use case error", func(t *testing.T) {
		mockUseCase := &MockAuthorize{
			FakeExecute: func(auth *domain.Auth) (*domain.AuthToken, error) {
				return nil, errors.New("unauthorized")
			},
		}
		controller := &AuthController{UseCase: mockUseCase}

		body, _ := json.Marshal(models.AuthRequest{ClientId: "id", ClientSecret: "secret"})
		req, _ := http.NewRequest(http.MethodPost, "/authorize", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		controller.Authorize(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("Expected status 401, got %v", rr.Code)
		}
	})
}

func TestAuthController_Validate(t *testing.T) {
	t.Run("should return 200 and token on success", func(t *testing.T) {
		mockUseCase := &MockAuthorize{
			FakeValidate: func(token string) (*domain.AuthToken, error) {
				return &domain.AuthToken{Token: token, ClientId: "test-client"}, nil
			},
		}
		controller := &AuthController{UseCase: mockUseCase}

		req, _ := http.NewRequest(http.MethodGet, "/validate?token=valid-token", nil)
		rr := httptest.NewRecorder()

		controller.Validate(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", rr.Code)
		}

		var resp models.AuthResponse
		json.Unmarshal(rr.Body.Bytes(), &resp)
		if resp.Token != "valid-token" {
			t.Errorf("Expected token valid-token, got %s", resp.Token)
		}
	})

	t.Run("should return 400 if token query param is missing", func(t *testing.T) {
		controller := &AuthController{UseCase: &MockAuthorize{}}
		req, _ := http.NewRequest(http.MethodGet, "/validate", nil)
		rr := httptest.NewRecorder()

		controller.Validate(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %v", rr.Code)
		}
	})

	t.Run("should return 401 on use case error", func(t *testing.T) {
		mockUseCase := &MockAuthorize{
			FakeValidate: func(token string) (*domain.AuthToken, error) {
				return nil, errors.New("invalid token")
			},
		}
		controller := &AuthController{UseCase: mockUseCase}

		req, _ := http.NewRequest(http.MethodGet, "/validate?token=invalid", nil)
		rr := httptest.NewRecorder()

		controller.Validate(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("Expected status 401, got %v", rr.Code)
		}
	})
}

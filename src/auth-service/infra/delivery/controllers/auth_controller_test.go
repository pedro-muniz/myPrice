package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	port "github.com/pedro-muniz/myPrice/auth/core/port/usecase/auth"
	"github.com/pedro-muniz/myPrice/auth/infra/delivery/models"
)

type MockAuthenticate struct {
	FakeExecute func(input *port.AuthenticateInput) (*port.AuthenticateOutput, error)
}

func (m *MockAuthenticate) Execute(input *port.AuthenticateInput) (*port.AuthenticateOutput, error) {
	if m.FakeExecute != nil {
		return m.FakeExecute(input)
	}
	return nil, nil
}

type MockAuthorize struct {
	FakeExecute func(token string) (*port.AuthorizeOutput, error)
}

func (m *MockAuthorize) Execute(token string) (*port.AuthorizeOutput, error) {
	if m.FakeExecute != nil {
		return m.FakeExecute(token)
	}
	return nil, nil
}

func TestAuthController_Authorize(t *testing.T) {
	t.Run("should return 200 and token on success", func(t *testing.T) {
		mockUseCase := &MockAuthenticate{
			FakeExecute: func(input *port.AuthenticateInput) (*port.AuthenticateOutput, error) {
				return &port.AuthenticateOutput{Token: "test-token", ExpiringIn: time.Hour}, nil
			},
		}
		controller := &AuthController{AuthenticateUseCase: mockUseCase}

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
		controller := &AuthController{AuthenticateUseCase: &MockAuthenticate{}}
		req, _ := http.NewRequest(http.MethodGet, "/authorize", nil)
		rr := httptest.NewRecorder()

		controller.Authorize(rr, req)

		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status 405, got %v", rr.Code)
		}
	})

	t.Run("should return 401 on use case error", func(t *testing.T) {
		mockUseCase := &MockAuthenticate{
			FakeExecute: func(input *port.AuthenticateInput) (*port.AuthenticateOutput, error) {
				return nil, errors.New("unauthorized")
			},
		}
		controller := &AuthController{AuthenticateUseCase: mockUseCase}

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
			FakeExecute: func(token string) (*port.AuthorizeOutput, error) {
				return &port.AuthorizeOutput{Token: token, ClientId: "test-client"}, nil
			},
		}
		controller := &AuthController{AuthorizeUseCase: mockUseCase}

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
		controller := &AuthController{AuthorizeUseCase: &MockAuthorize{}}
		req, _ := http.NewRequest(http.MethodGet, "/validate", nil)
		rr := httptest.NewRecorder()

		controller.Validate(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %v", rr.Code)
		}
	})

	t.Run("should return 401 on use case error", func(t *testing.T) {
		mockUseCase := &MockAuthorize{
			FakeExecute: func(token string) (*port.AuthorizeOutput, error) {
				return nil, errors.New("invalid token")
			},
		}
		controller := &AuthController{AuthorizeUseCase: mockUseCase}

		req, _ := http.NewRequest(http.MethodGet, "/validate?token=invalid", nil)
		rr := httptest.NewRecorder()

		controller.Validate(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("Expected status 401, got %v", rr.Code)
		}
	})
}

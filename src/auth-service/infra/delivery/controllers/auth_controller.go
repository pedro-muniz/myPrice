package controllers

import (
	"encoding/json"
	"net/http"

	port "github.com/pedro-muniz/myPrice/auth/core/port/usecase/auth"
	"github.com/pedro-muniz/myPrice/auth/infra/delivery/converters"
	"github.com/pedro-muniz/myPrice/auth/infra/delivery/models"
)

type AuthController struct {
	UseCase port.IAuthorize
}

// Authorize handles client authentication and token generation
func (this *AuthController) Authorize(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	domainAuth := converters.ToDomain(request)
	authToken, err := this.UseCase.Execute(domainAuth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := converters.ToResponse(authToken)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Validate handles JWT token verification and expiration check
func (this *AuthController) Validate(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	authToken, err := this.UseCase.Validate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := converters.ToResponse(authToken)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

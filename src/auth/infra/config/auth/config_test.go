package auth

import "testing"

func TestCreateAuthorizeUseCase_shouldPass(t *testing.T) {
	authConfig := &AuthConfig{}
	authConfig.CreateAuthorizeUseCase()
}

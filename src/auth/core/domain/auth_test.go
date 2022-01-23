package domain

import (
	authErrors "core/customErrors/auth"
	"testing"
)

//test GenerateToken with an invalid auth object
func TestGenerateToken_invalidClientId_shouldReturnAnError(t *testing.T) {
	//arrange
	expectedError := authErrors.InvalidClientId("")
	auth := &Auth{}

	//act
	token, authError := auth.GenerateToken()

	//assert
	if len(token) > 0 {
		t.Error("The GenerateToken returned a strange token")
	}

	if authError == nil {
		t.Error("The GenerateToken didn't return the expectedError")
	}

	if expectedError != authError {
		t.Error("The GenerateToken didn't return the expectedError")
	}
}

func TestGenerateToken_validClientId_shouldReturnToken(t *testing.T) {
	//arrange
	auth := &Auth{}
	auth.ClientId = "testing"

	//act
	token, authError := auth.GenerateToken()

	//assert
	if authError != nil {
		t.Errorf("Error generating token %s", authError.Error())
	}

	if len(token) <= 0 {
		t.Error("The GenerateToken didn't return the token")
	}

	t.Log(token)
}

package domain

import (
	"testing"

	authErrors "github.com/pedro-muniz/myPrice/auth/core/customerror/auth"
)

//test GenerateToken with an invalid auth object
func TestGenerateToken_invalidClientId_shouldReturnAnError(t *testing.T) {
	//arrange
	expectedError := authErrors.InvalidClientId("")
	auth := &Auth{}

	//act
	token, authError := auth.generateToken()

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
	token, authError := auth.generateToken()

	//assert
	if authError != nil {
		t.Errorf("Error generating token %s", authError.Error())
	}

	if len(token) <= 0 {
		t.Error("The GenerateToken didn't return the token")
	}

	t.Log(token)
}

func TestGenerateAuthToken_validData_shouldReturnAuthTokenStruct(t *testing.T) {
	//arrange
	auth := &Auth{}
	auth.ClientId = "testing"

	//act
	authToken, err := auth.GetAuthToken()

	//assert
	if err != nil {
		t.Errorf("Error generating token %s", err.Error())
	}

	if authToken == nil {
		t.Errorf("Error generating token, the method didn't return the token structure")
	}

	if len(authToken.ClientId) <= 0 {
		t.Errorf("Error generating token, the clientId is invalid")
	}

	if authToken.ExpiringAt.IsZero() {
		t.Errorf("Error generating token, invalid expiring date")
	}

	if len(authToken.Token) <= 0 {
		t.Errorf("Error generating token, invalid token")
	}

}

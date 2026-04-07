package domain

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authErrors "github.com/pedro-muniz/myPrice/auth/core/customerror/auth"
)

// test GenerateToken with an invalid auth object
func TestGenerateToken_invalidClientId_shouldReturnAnError(t *testing.T) {
	//arrange
	expectedError := authErrors.InvalidClientId("")
	auth := &Auth{}

	//act
	token, authError := auth.generateToken(time.Hour)

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
	token, authError := auth.generateToken(time.Hour)

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

	if authToken.ExpiringIn <= time.Hour {
		t.Errorf("Error generating token, invalid expiring date %s", authToken.ExpiringIn)
	}

	if len(authToken.Token) <= 0 {
		t.Errorf("Error generating token, invalid token")
	}

}

func TestGetAuthToken_shouldReturnValidJWT(t *testing.T) {
	// arrange
	secret := "test_secret"
	t.Setenv("auth_secret", secret)

	auth := &Auth{
		ClientId: "test_client",
	}

	// act
	authToken, err := auth.GetAuthToken()

	// assert
	if err != nil {
		t.Fatalf("Failed to get auth token: %v", err)
	}

	// Parsing and validating the JWT
	parsedToken, err := jwt.Parse(authToken.Token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		t.Fatalf("Failed to parse JWT: %v", err)
	}

	if !parsedToken.Valid {
		t.Error("Token is not valid")
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		if claims["client_id"] != "test_client" {
			t.Errorf("Expected client_id test_client, got %v", claims["client_id"])
		}
	} else {
		t.Error("Could not parse claims")
	}
}

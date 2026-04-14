package authorize

import (
	"errors"
	"testing"

	authErrors "github.com/pedro-muniz/myPrice/auth/core/customerror/auth"
	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
	port "github.com/pedro-muniz/myPrice/auth/core/port/usecase/auth"
)

func TestAuthenticateExecute_invalidAuthObject_shouldReturnAnError(t *testing.T) {
	//arrange
	expectedError := authErrors.InvalidAuthReference("")
	authenticateUseCase := &Authenticate{AuthRepository: mockedRepo, AuthPublisher: mockedPub}

	//act
	authToken, authError := authenticateUseCase.Execute(nil)

	//assert
	if authToken != nil {
		t.Error("The authenticate returned a strange token")
	}

	if authError == nil {
		t.Error("The authenticate didn't return the expectedError")
	}

	if expectedError != authError {
		t.Error("The authenticate didn't return the expectedError")
	}

}

func TestAuthenticateExecute_errorReturningDatabaseRecord_shouldReturnAnError(t *testing.T) {
	//arrange
	err := errors.New("error")
	expectedError := authErrors.ErrorGettingAuthDatabaseRecord(err.Error())
	input := &port.AuthenticateInput{ClientId: "test", ClientSecret: "Test"}
	repository := &TestAuthRepository{
		FakeGet: func(email string, password string) (<-chan *domain.Auth, <-chan error) {
			ec := make(chan error, 1)
			ec <- err
			return nil, ec
		},
	}

	authenticateUseCase := &Authenticate{AuthRepository: repository, AuthPublisher: mockedPub}

	//act
	authToken, authError := authenticateUseCase.Execute(input)

	//assert
	if authToken != nil {
		t.Error("The authenticate returned a strange token")
	}

	if expectedError != authError {
		t.Error("The authenticate didn't return the expectedError")
	}
}

func TestAuthenticateExecute_clientNotFound_shouldReturnAnError(t *testing.T) {
	//arrange
	expectedError := authErrors.ClientNotFound("")
	input := &port.AuthenticateInput{ClientId: "test", ClientSecret: "Test"}
	repository := &TestAuthRepository{
		FakeGet: func(email string, password string) (<-chan *domain.Auth, <-chan error) {
			ac := make(chan *domain.Auth, 1)
			close(ac)
			return ac, nil
		},
	}

	authenticateUseCase := &Authenticate{AuthRepository: repository, AuthPublisher: mockedPub}

	//act
	authToken, authError := authenticateUseCase.Execute(input)

	//assert
	if authToken != nil {
		t.Error("The authenticate returned a strange token")
	}

	if expectedError != authError {
		t.Error("The authenticate didn't return the expectedError")
	}
}

func TestAuthenticateExecute_dataOk_shouldReturnAToken(t *testing.T) {
	//arrange
	input := &port.AuthenticateInput{ClientId: "test", ClientSecret: "Test"}
	auth := &domain.Auth{ClientId: "test", ClientSecret: "Test"}
	repository := &TestAuthRepository{
		FakeGet: func(email string, password string) (<-chan *domain.Auth, <-chan error) {
			ac := make(chan *domain.Auth, 1)
			ac <- auth
			return ac, nil
		},
	}

	authenticateUseCase := &Authenticate{AuthRepository: repository, AuthPublisher: mockedPub}

	//act
	authToken, _ := authenticateUseCase.Execute(input)

	//assert
	if authToken == nil || len(authToken.Token) <= 0 {
		t.Error("The authenticate didn't return the token")
	}

	if len(authToken.ClientId) <= 0 {
		t.Error("The authenticate didn't return the clientId")
	}

	if authToken.ExpiringIn <= 0 {
		t.Error("The authenticate didn't return the ExpiringIn")
	}

	t.Log(authToken.Token)
}

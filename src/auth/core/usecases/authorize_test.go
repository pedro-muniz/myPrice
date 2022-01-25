package usecases

import (
	authErrors "core/customErrors/auth"
	domain "core/domain"
	"errors"
	"testing"
)

type MockAuthRepository struct {
	FakeGet  func(clientId string, clientSecret string) (*domain.Auth, error)
	FakeSave func(auth *domain.Auth) error
}

func (repository *MockAuthRepository) Get(clientId string, clientSecret string) (*domain.Auth, error) {
	if repository.FakeGet != nil {
		return repository.FakeGet(clientId, clientSecret)
	}

	return nil, nil
}

func (repository *MockAuthRepository) Save(auth *domain.Auth) error {
	if repository.FakeSave != nil {
		return repository.FakeSave(auth)
	}

	return nil
}

func TestUseCaseExecute_invalidAuthObject_shouldReturnAnError(t *testing.T) {
	//arrange
	expectedError := authErrors.InvalidAuthReference("")
	repository := &MockAuthRepository{}
	authorizeUseCase := &Authorize{Auth: nil, AuthRepository: repository}

	//act
	token, authError := authorizeUseCase.Execute()

	//assert
	if len(token) > 0 {
		t.Error("The authorize returned a strange token")
	}

	if authError == nil {
		t.Error("The authorize didn't return the expectedError")
	}

	if expectedError != authError {
		t.Error("The authorize didn't return the expectedError")
	}

}

func TestUseCaseExecute_errorReturningDatabaseRecord_shouldReturnAnError(t *testing.T) {
	//arrange
	err := errors.New("error")
	expectedError := authErrors.ErrorGettingAuthDatabaseRecord(err.Error())
	auth := &domain.Auth{ClientId: "test", ClientSecret: "Test"}
	repository := &MockAuthRepository{
		FakeGet: func(clientId string, clientSecret string) (*domain.Auth, error) {
			return nil, err
		},
	}

	authorizeUseCase := &Authorize{Auth: auth, AuthRepository: repository}

	//act
	token, authError := authorizeUseCase.Execute()

	//assert
	if len(token) > 0 {
		t.Error("The authorize returned a strange token")
	}

	if expectedError != authError {
		t.Error("The authorize didn't return the expectedError")
	}
}

func TestUseCaseExecute_clientNotFound_shouldReturnAnError(t *testing.T) {
	//arrange
	expectedError := authErrors.ClientNotFound("")
	auth := &domain.Auth{ClientId: "test", ClientSecret: "Test"}
	repository := &MockAuthRepository{
		FakeGet: func(clientId string, clientSecret string) (*domain.Auth, error) {
			return nil, nil
		},
	}

	authorizeUseCase := &Authorize{Auth: auth, AuthRepository: repository}

	//act
	token, authError := authorizeUseCase.Execute()

	//assert
	if len(token) > 0 {
		t.Error("The authorize returned a strange token")
	}

	if expectedError != authError {
		t.Error("The authorize didn't return the expectedError")
	}
}

func TestUseCaseExecute_dataOk_shouldReturnAToken(t *testing.T) {
	//arrange
	auth := &domain.Auth{ClientId: "test", ClientSecret: "Test"}
	repository := &MockAuthRepository{
		FakeGet: func(clientId string, clientSecret string) (*domain.Auth, error) {
			return auth, nil
		},
	}

	authorizeUseCase := &Authorize{Auth: auth, AuthRepository: repository}

	//act
	token, _ := authorizeUseCase.Execute()

	//assert
	if len(token) <= 0 {
		t.Error("The authorize didn't return the token")
	}

	t.Log(token)
}

package authorize

import (
	"errors"
	"testing"
	"time"

	domain "github.com/pedro-muniz/myPrice/auth/core/domain"

	authErrors "github.com/pedro-muniz/myPrice/auth/core/customerror/auth"
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

type MockAuthPublisher struct {
	FakePublish func(token string, expiringAt time.Duration) error
	FakeGet     func(token string) (*domain.AuthToken, error)
}

func (publisher *MockAuthPublisher) Publish(token string, expiringAt time.Duration) error {
	if publisher.FakePublish != nil {
		return publisher.FakePublish(token, expiringAt)
	}

	return nil
}

func (publisher *MockAuthPublisher) Get(token string) (*domain.AuthToken, error) {
	if publisher.FakeGet != nil {
		return publisher.FakeGet(token)
	}

	return nil, nil
}

var mockedRepo *MockAuthRepository = &MockAuthRepository{}
var mockedPub *MockAuthPublisher = &MockAuthPublisher{}

func TestUseCaseExecute_invalidAuthObject_shouldReturnAnError(t *testing.T) {
	//arrange
	expectedError := authErrors.InvalidAuthReference("")
	authorizeUseCase := &Authorize{AuthRepository: mockedRepo, AuthPublisher: mockedPub}

	//act
	authToken, authError := authorizeUseCase.Execute(nil)

	//assert
	if authToken != nil {
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

	authorizeUseCase := &Authorize{AuthRepository: repository, AuthPublisher: mockedPub}

	//act
	authToken, authError := authorizeUseCase.Execute(auth)

	//assert
	if authToken != nil {
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

	authorizeUseCase := &Authorize{AuthRepository: repository, AuthPublisher: mockedPub}

	//act
	authToken, authError := authorizeUseCase.Execute(auth)

	//assert
	if authToken != nil {
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

	authorizeUseCase := &Authorize{AuthRepository: repository, AuthPublisher: mockedPub}

	//act
	authToken, _ := authorizeUseCase.Execute(auth)

	//assert
	if authToken == nil || len(authToken.Token) <= 0 {
		t.Error("The authorize didn't return the token")
	}

	if len(authToken.ClientId) <= 0 {
		t.Error("The authorize didn't return the clientId")
	}

	if authToken.ExpiringAt <= 0 {
		t.Error("The authorize didn't return the ExpiringAt")
	}

	t.Log(authToken.Token)
}

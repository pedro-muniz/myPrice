package usecases

import (
	authErrors "core/customErrors/auth"
	domain "core/domain"
	"testing"
)

type MockRepository struct {
	FakeGet  func(clientId string, clientSecret string) (*domain.Auth, error)
	FakeSave func(auth *domain.Auth) error
}

func (repository *MockRepository) Get(clientId string, clientSecret string) (*domain.Auth, error) {
	if repository.FakeGet != nil {
		return repository.FakeGet(clientId, clientSecret)
	}

	return nil, nil
}

func (repository *MockRepository) Save(auth *domain.Auth) error {
	if repository.FakeSave != nil {
		return repository.FakeSave(auth)
	}

	return nil
}

func TestUseCaseExecute_invalidAuthObject_shouldReturnAnError(t *testing.T) {
	//arrange
	expectedError := authErrors.InvalidAuthReference("")
	repository := &MockRepository{}
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

package authorize

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authErrors "github.com/pedro-muniz/myPrice/auth/core/customerror/auth"
	domain "github.com/pedro-muniz/myPrice/auth/core/domain"
	repository "github.com/pedro-muniz/myPrice/auth/core/port/repository"
)

var _ repository.IAuthRepository = (*TestAuthRepository)(nil)

type TestAuthRepository struct {
	FakeGet  func(string, string) (<-chan *domain.Auth, <-chan error)
	FakeSave func(*domain.Auth) <-chan error
}

func (m *TestAuthRepository) Get(email string, password string) (<-chan *domain.Auth, <-chan error) {
	if m.FakeGet != nil {
		return m.FakeGet(email, password)
	}

	ac := make(chan *domain.Auth, 1)
	ec := make(chan error, 1)
	close(ac)
	close(ec)
	return ac, ec
}

func (m *TestAuthRepository) Save(auth *domain.Auth) <-chan error {
	if m.FakeSave != nil {
		return m.FakeSave(auth)
	}

	ec := make(chan error, 1)
	close(ec)
	return ec
}

type MockAuthPublisher struct {
	FakePublish func(token string, expiringIn time.Duration) error
	FakeGet     func(token string) (*domain.AuthToken, error)
}

func (publisher *MockAuthPublisher) Publish(token string, expiringIn time.Duration) error {
	if publisher.FakePublish != nil {
		return publisher.FakePublish(token, expiringIn)
	}

	return nil
}

func (publisher *MockAuthPublisher) Get(token string) (*domain.AuthToken, error) {
	if publisher.FakeGet != nil {
		return publisher.FakeGet(token)
	}

	return nil, nil
}

var mockedRepo *TestAuthRepository = &TestAuthRepository{}
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
	repository := &TestAuthRepository{
		FakeGet: func(email string, password string) (<-chan *domain.Auth, <-chan error) {
			ec := make(chan error, 1)
			ec <- err
			return nil, ec
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
	repository := &TestAuthRepository{
		FakeGet: func(email string, password string) (<-chan *domain.Auth, <-chan error) {
			ac := make(chan *domain.Auth, 1)
			close(ac)
			return ac, nil
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
	repository := &TestAuthRepository{
		FakeGet: func(email string, password string) (<-chan *domain.Auth, <-chan error) {
			ac := make(chan *domain.Auth, 1)
			ac <- auth
			return ac, nil
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

	if authToken.ExpiringIn <= 0 {
		t.Error("The authorize didn't return the ExpiringIn")
	}

	t.Log(authToken.Token)
}

func TestUseCaseValidate_tokenNotFoundInRedis_shouldReturnAnError(t *testing.T) {
	// arrange
	publisher := &MockAuthPublisher{
		FakeGet: func(token string) (*domain.AuthToken, error) {
			return nil, nil // Not found
		},
	}
	authorizeUseCase := &Authorize{AuthRepository: mockedRepo, AuthPublisher: publisher}

	// act
	authToken, authError := authorizeUseCase.Validate("any-token")

	// assert
	if authToken != nil {
		t.Error("The validate returned a token when it should not")
	}
	if authError == nil {
		t.Error("The validate didn't return an error")
	}
}

func TestUseCaseValidate_invalidJWT_shouldReturnAnError(t *testing.T) {
	// arrange
	publisher := &MockAuthPublisher{
		FakeGet: func(token string) (*domain.AuthToken, error) {
			return &domain.AuthToken{Token: token}, nil
		},
	}
	authorizeUseCase := &Authorize{AuthRepository: mockedRepo, AuthPublisher: publisher}

	// act
	authToken, authError := authorizeUseCase.Validate("not-a-jwt")

	// assert
	if authToken != nil {
		t.Error("The validate returned a token for invalid JWT")
	}
	if authError == nil {
		t.Error("The validate didn't return an error for invalid JWT")
	}
}

func TestUseCaseValidate_validToken_shouldReturnAuthToken(t *testing.T) {
	// arrange
	secret := "test_secret"
	t.Setenv("auth_secret", secret)
	clientId := "test-client"

	// Create a valid JWT
	claims := jwt.MapClaims{
		"client_id": clientId,
		"exp":       time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret))

	publisher := &MockAuthPublisher{
		FakeGet: func(token string) (*domain.AuthToken, error) {
			return &domain.AuthToken{Token: token}, nil
		},
	}
	authorizeUseCase := &Authorize{AuthRepository: mockedRepo, AuthPublisher: publisher}

	// act
	authToken, authError := authorizeUseCase.Validate(tokenString)

	// assert
	if authError != nil {
		t.Errorf("The validate returned an error: %v", authError)
	}
	if authToken == nil {
		t.Fatal("The validate didn't return a token")
	}
	if authToken.ClientId != clientId {
		t.Errorf("Expected client id %s, got %s", clientId, authToken.ClientId)
	}
}

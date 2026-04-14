package authorize

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func TestAuthorizeExecute_tokenNotFoundInRedis_shouldReturnAnError(t *testing.T) {
	// arrange
	publisher := &MockAuthPublisher{
		FakeGet: func(token string) (*domain.AuthToken, error) {
			return nil, nil // Not found
		},
	}
	authorizeUseCase := &Authorize{AuthPublisher: publisher}

	// act
	authToken, authError := authorizeUseCase.Execute("any-token")

	// assert
	if authToken != nil {
		t.Error("The authorize returned a token when it should not")
	}
	if authError == nil {
		t.Error("The authorize didn't return an error")
	}
}

func TestAuthorizeExecute_invalidJWT_shouldReturnAnError(t *testing.T) {
	// arrange
	publisher := &MockAuthPublisher{
		FakeGet: func(token string) (*domain.AuthToken, error) {
			return &domain.AuthToken{Token: token}, nil
		},
	}
	authorizeUseCase := &Authorize{AuthPublisher: publisher}

	// act
	authToken, authError := authorizeUseCase.Execute("not-a-jwt")

	// assert
	if authToken != nil {
		t.Error("The authorize returned a token for invalid JWT")
	}
	if authError == nil {
		t.Error("The authorize didn't return an error for invalid JWT")
	}
}

func TestAuthorizeExecute_validToken_shouldReturnAuthToken(t *testing.T) {
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
	authorizeUseCase := &Authorize{AuthPublisher: publisher}

	// act
	authToken, authError := authorizeUseCase.Execute(tokenString)

	// assert
	if authError != nil {
		t.Errorf("The authorize returned an error: %v", authError)
	}
	if authToken == nil {
		t.Fatal("The authorize didn't return a token")
	}
	if authToken.ClientId != clientId {
		t.Errorf("Expected client id %s, got %s", clientId, authToken.ClientId)
	}
}

package product

import (
	"errors"
	"testing"

	productErrors "github.com/pedro-muniz/myPrice/src/product-service/core/customerror/product"
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	port "github.com/pedro-muniz/myPrice/src/product-service/core/port/usecase/product"
)

func TestProductManagementCreate_nilInput_shouldReturnError(t *testing.T) {
	uc := &ProductManagement{}
	_, err := uc.Create(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestProductManagementCreate_repoError_shouldReturnError(t *testing.T) {
	input := &port.CreateInput{Name: "Test"}
	repoErr := errors.New("db error")
	repository := &TestProductRepository{
		FakeSave: func(product *domain.Product) (<-chan *domain.Product, <-chan error) {
			ec := make(chan error, 1)
			ec <- repoErr
			return nil, ec
		},
	}
	uc := &ProductManagement{ProductRepository: repository}

	_, err := uc.Create(input)

	if err == nil {
		t.Error("expected error")
	}
	expectedError := productErrors.ErrorSavingProductDatabaseRecord(repoErr.Error())
	if err.Error() != expectedError.Error() {
		t.Errorf("expected %v, got %v", expectedError, err)
	}
}

func TestProductManagementCreate_success_shouldReturnId(t *testing.T) {
	input := &port.CreateInput{Name: "Test"}
	savedProduct := &domain.Product{Id: "123"}
	repository := &TestProductRepository{
		FakeSave: func(product *domain.Product) (<-chan *domain.Product, <-chan error) {
			pc := make(chan *domain.Product, 1)
			pc <- savedProduct
			return pc, nil
		},
	}
	uc := &ProductManagement{ProductRepository: repository}

	output, err := uc.Create(input)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if output.Id != "123" {
		t.Errorf("expected id 123, got %s", output.Id)
	}
}

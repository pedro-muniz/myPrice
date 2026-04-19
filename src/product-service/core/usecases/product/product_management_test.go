package product

import (
	"errors"
	"testing"
	"time"

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

func TestProductManagementCreate_invalidInput_shouldReturnError(t *testing.T) {
	input := &port.CreateInput{Name: ""} // Invalid: missing barcode, etc.
	uc := &ProductManagement{}

	_, err := uc.Create(input)

	if err == nil {
		t.Error("expected error for invalid input")
	}
}

func TestProductManagementCreate_repoError_shouldReturnError(t *testing.T) {
	input := &port.CreateInput{
		BarCode:     "123",
		Name:        "Test",
		Description: "Desc",
		Gross:       10,
		Net:         10,
		Selling:     10,
		Recommended: 10,
	}
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
	input := &port.CreateInput{
		BarCode:     "123",
		Name:        "Test",
		Description: "Desc",
		Gross:       10,
		Net:         10,
		Selling:     10,
		Recommended: 10,
	}
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
		t.Fatalf("expected no error, got %v", err)
	}
	if output.Id != "123" {
		t.Errorf("expected id 123, got %s", output.Id)
	}
}

func TestProductManagementUpdate_success_shouldReturnSuccess(t *testing.T) {
	input := &port.UpdateInput{
		Id:          "123",
		BarCode:     "123",
		Name:        "Test",
		Description: "Desc",
		Gross:       10,
		Net:         10,
		Selling:     10,
		Recommended: 10,
	}
	existingProduct := &domain.Product{
		Id:        "123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Price: domain.Price{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	repository := &TestProductRepository{
		FakeFindById: func(id string) (<-chan *domain.Product, <-chan error) {
			pc := make(chan *domain.Product, 1)
			pc <- existingProduct
			return pc, nil
		},
		FakeUpdate: func(product *domain.Product) <-chan error {
			ec := make(chan error, 1)
			ec <- nil
			return ec
		},
	}
	uc := &ProductManagement{ProductRepository: repository}

	output, err := uc.Update(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !output.Success {
		t.Error("expected success true")
	}
}

func TestProductManagementUpdate_invalidInput_shouldReturnError(t *testing.T) {
	input := &port.UpdateInput{
		Id:   "123",
		Name: "", // Invalid
	}
	existingProduct := &domain.Product{
		Id:        "123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Price: domain.Price{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	repository := &TestProductRepository{
		FakeFindById: func(id string) (<-chan *domain.Product, <-chan error) {
			pc := make(chan *domain.Product, 1)
			pc <- existingProduct
			return pc, nil
		},
	}
	uc := &ProductManagement{ProductRepository: repository}

	_, err := uc.Update(input)

	if err == nil {
		t.Error("expected error for invalid input")
	}
}

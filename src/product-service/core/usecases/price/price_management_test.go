package price

import (
	"errors"
	"testing"
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	port "github.com/pedro-muniz/myPrice/src/product-service/core/port/usecase/price"
)

func TestPriceManagementCreate_nilInput_shouldReturnError(t *testing.T) {
	uc := &PriceManagement{}
	_, err := uc.Create(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestPriceManagementCreate_invalidInput_shouldReturnError(t *testing.T) {
	input := &port.CreateInput{Gross: -1} // Invalid
	uc := &PriceManagement{}

	_, err := uc.Create(input)

	if err == nil {
		t.Error("expected error for invalid input")
	}
}

func TestPriceManagementCreate_success_shouldReturnId(t *testing.T) {
	input := &port.CreateInput{
		CompanyId:   "c1",
		BranchId:    "b1",
		Gross:       100,
		Net:         80,
		Selling:     120,
		Recommended: 150,
	}
	savedPrice := &domain.Price{Id: "price-123"}
	repository := &TestPriceRepository{
		FakeSave: func(price *domain.Price) (<-chan *domain.Price, <-chan error) {
			pc := make(chan *domain.Price, 1)
			pc <- savedPrice
			return pc, nil
		},
	}
	uc := &PriceManagement{PriceRepository: repository}

	output, err := uc.Create(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if output.Id != "price-123" {
		t.Errorf("expected id price-123, got %s", output.Id)
	}
}

func TestPriceManagementUpdate_success_shouldReturnSuccess(t *testing.T) {
	input := &port.UpdateInput{
		CompanyId:   "c1",
		BranchId:    "b1",
		Id:          "123",
		Gross:       100,
		Net:         80,
		Selling:     120,
		Recommended: 150,
	}
	existingPrice := &domain.Price{
		CompanyId: "c1",
		BranchId:  "b1",
		Id:        "123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	repository := &TestPriceRepository{
		FakeFindById: func(companyId string, branchId string, id string) (<-chan *domain.Price, <-chan error) {
			pc := make(chan *domain.Price, 1)
			pc <- existingPrice
			return pc, nil
		},
		FakeUpdate: func(price *domain.Price) <-chan error {
			ec := make(chan error, 1)
			ec <- nil
			return ec
		},
	}
	uc := &PriceManagement{PriceRepository: repository}

	output, err := uc.Update(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !output.Success {
		t.Error("expected success true")
	}
}

func TestPriceManagementUpdate_invalidInput_shouldReturnError(t *testing.T) {
	input := &port.UpdateInput{
		CompanyId: "c1",
		BranchId:  "b1",
		Id:        "123",
		Gross:     -1, // Invalid
	}
	existingPrice := &domain.Price{
		CompanyId: "c1",
		BranchId:  "b1",
		Id:        "123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	repository := &TestPriceRepository{
		FakeFindById: func(companyId string, branchId string, id string) (<-chan *domain.Price, <-chan error) {
			pc := make(chan *domain.Price, 1)
			pc <- existingPrice
			return pc, nil
		},
	}
	uc := &PriceManagement{PriceRepository: repository}

	_, err := uc.Update(input)

	if err == nil {
		t.Error("expected error for invalid input")
	}
}

func TestPriceManagementDelete_nilInput_shouldReturnError(t *testing.T) {
	uc := &PriceManagement{}
	_, err := uc.Delete(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestPriceManagementDelete_emptyId_shouldReturnError(t *testing.T) {
	input := &port.DeleteInput{
		CompanyId: "c1",
		BranchId:  "b1",
		Id:        "",
	}
	uc := &PriceManagement{}

	_, err := uc.Delete(input)

	if err == nil {
		t.Error("expected error for empty id")
	}
}

func TestPriceManagementDelete_repoError_shouldReturnError(t *testing.T) {
	input := &port.DeleteInput{
		CompanyId: "c1",
		BranchId:  "b1",
		Id:        "123",
	}
	repository := &TestPriceRepository{
		FakeDelete: func(companyId string, branchId string, id string, deletedAt time.Time) <-chan error {
			ec := make(chan error, 1)
			ec <- errors.New("db error")
			return ec
		},
	}
	uc := &PriceManagement{PriceRepository: repository}

	_, err := uc.Delete(input)

	if err == nil {
		t.Error("expected error from repository")
	}
}

func TestPriceManagementDelete_success_shouldReturnSuccess(t *testing.T) {
	input := &port.DeleteInput{
		CompanyId: "c1",
		BranchId:  "b1",
		Id:        "123",
	}
	repository := &TestPriceRepository{
		FakeDelete: func(companyId string, branchId string, id string, deletedAt time.Time) <-chan error {
			ec := make(chan error, 1)
			ec <- nil
			return ec
		},
	}
	uc := &PriceManagement{PriceRepository: repository}

	output, err := uc.Delete(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !output.Success {
		t.Error("expected success true")
	}
}

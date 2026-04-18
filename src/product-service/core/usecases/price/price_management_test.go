package price

import (
	"testing"

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

func TestPriceManagementCreate_success_shouldReturnId(t *testing.T) {
	input := &port.CreateInput{Gross: 100}
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
		t.Errorf("expected no error, got %v", err)
	}
	if output.Id != "price-123" {
		t.Errorf("expected id price-123, got %s", output.Id)
	}
}

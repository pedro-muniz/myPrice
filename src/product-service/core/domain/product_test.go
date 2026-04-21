package domain

import (
	"testing"
	"time"
)

func TestNewProduct_shouldSetFields(t *testing.T) {
	price := Price{Gross: 10, Net: 8, Selling: 12, Recommended: 15}
	product := NewProduct("c1", "b1", "bar123", "Test", "Desc", 5.0, price, 20.0, National)

	if product.CompanyId != "c1" {
		t.Errorf("expected CompanyId c1, got %s", product.CompanyId)
	}
	if product.BranchId != "b1" {
		t.Errorf("expected BranchId b1, got %s", product.BranchId)
	}
	if product.BarCode != "bar123" {
		t.Errorf("expected BarCode bar123, got %s", product.BarCode)
	}
	if product.Name != "Test" {
		t.Errorf("expected Name Test, got %s", product.Name)
	}
	if product.Description != "Desc" {
		t.Errorf("expected Description Desc, got %s", product.Description)
	}
	if product.Discount != 5.0 {
		t.Errorf("expected Discount 5.0, got %f", product.Discount)
	}
	if product.Price.Gross != 10 {
		t.Errorf("expected Price.Gross 10, got %f", product.Price.Gross)
	}
	if product.ProfitPercent != 20.0 {
		t.Errorf("expected ProfitPercent 20.0, got %f", product.ProfitPercent)
	}
	if product.Origin != National {
		t.Errorf("expected Origin National, got %d", product.Origin)
	}
	if product.Id != "" {
		t.Errorf("expected empty Id, got %s", product.Id)
	}
}

func validProduct() *Product {
	now := time.Now()
	return &Product{
		CompanyId:     "c1",
		BranchId:      "b1",
		BarCode:       "bar123",
		Name:          "Test",
		Description:   "Desc",
		Discount:      5.0,
		ProfitPercent: 20.0,
		Origin:        National,
		Price: Price{
			CompanyId:   "c1",
			BranchId:    "b1",
			Gross:       10,
			Net:         8,
			Selling:     12,
			Recommended: 15,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func TestProductValidate_valid_shouldReturnNil(t *testing.T) {
	product := validProduct()
	if err := product.Validate(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestProductValidate_emptyCompanyId_shouldReturnError(t *testing.T) {
	product := validProduct()
	product.CompanyId = ""
	err := product.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid product company id" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestProductValidate_emptyBranchId_shouldReturnError(t *testing.T) {
	product := validProduct()
	product.BranchId = ""
	err := product.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid product branch id" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestProductValidate_emptyBarCode_shouldReturnError(t *testing.T) {
	product := validProduct()
	product.BarCode = ""
	err := product.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid product barcode" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestProductValidate_emptyName_shouldReturnError(t *testing.T) {
	product := validProduct()
	product.Name = ""
	err := product.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid product name" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestProductValidate_emptyDescription_shouldReturnError(t *testing.T) {
	product := validProduct()
	product.Description = ""
	err := product.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid product description" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestProductValidate_invalidOrigin_shouldReturnError(t *testing.T) {
	product := validProduct()
	product.Origin = Origin(99)
	err := product.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid product origin" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestProductValidate_invalidPrice_shouldReturnError(t *testing.T) {
	product := validProduct()
	product.Price.Gross = 0 // Make embedded price invalid
	err := product.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid product price" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestProductValidate_zeroCreatedAt_shouldReturnError(t *testing.T) {
	product := validProduct()
	product.CreatedAt = time.Time{}
	err := product.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid product created at" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestProductValidate_zeroUpdatedAt_shouldReturnError(t *testing.T) {
	product := validProduct()
	product.UpdatedAt = time.Time{}
	err := product.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid product updated at" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestProductDelete_shouldSetDeletedAtOnProductAndPrice(t *testing.T) {
	product := validProduct()
	if product.DeletedAt != nil {
		t.Fatal("expected DeletedAt to be nil before Delete")
	}
	if product.Price.DeletedAt != nil {
		t.Fatal("expected Price.DeletedAt to be nil before Delete")
	}

	before := time.Now()
	product.Delete()
	after := time.Now()

	if product.DeletedAt == nil {
		t.Fatal("expected DeletedAt to be set after Delete")
	}
	if product.DeletedAt.Before(before) || product.DeletedAt.After(after) {
		t.Error("DeletedAt should be between before and after timestamps")
	}
	if product.Price.DeletedAt == nil {
		t.Fatal("expected Price.DeletedAt to be set after Delete")
	}
}

package domain

import (
	"testing"
	"time"
)

func TestNewPrice_shouldSetFields(t *testing.T) {
	price := NewPrice("c1", "b1", 10.0, 8.0, 12.0, 15.0)

	if price.CompanyId != "c1" {
		t.Errorf("expected CompanyId c1, got %s", price.CompanyId)
	}
	if price.BranchId != "b1" {
		t.Errorf("expected BranchId b1, got %s", price.BranchId)
	}
	if price.Gross != 10.0 {
		t.Errorf("expected Gross 10.0, got %f", price.Gross)
	}
	if price.Net != 8.0 {
		t.Errorf("expected Net 8.0, got %f", price.Net)
	}
	if price.Selling != 12.0 {
		t.Errorf("expected Selling 12.0, got %f", price.Selling)
	}
	if price.Recommended != 15.0 {
		t.Errorf("expected Recommended 15.0, got %f", price.Recommended)
	}
	if price.Id != "" {
		t.Errorf("expected empty Id, got %s", price.Id)
	}
}

func validPrice() *Price {
	return &Price{
		CompanyId:   "c1",
		BranchId:    "b1",
		Gross:       10.0,
		Net:         8.0,
		Selling:     12.0,
		Recommended: 15.0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func TestPriceValidate_valid_shouldReturnNil(t *testing.T) {
	price := validPrice()
	if err := price.Validate(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestPriceValidate_emptyCompanyId_shouldReturnError(t *testing.T) {
	price := validPrice()
	price.CompanyId = ""
	err := price.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid price company id" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPriceValidate_emptyBranchId_shouldReturnError(t *testing.T) {
	price := validPrice()
	price.BranchId = ""
	err := price.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid price branch id" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPriceValidate_invalidGross_shouldReturnError(t *testing.T) {
	price := validPrice()
	price.Gross = 0
	err := price.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid price gross" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPriceValidate_invalidNet_shouldReturnError(t *testing.T) {
	price := validPrice()
	price.Net = -1
	err := price.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid price net" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPriceValidate_invalidSelling_shouldReturnError(t *testing.T) {
	price := validPrice()
	price.Selling = 0
	err := price.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid price selling" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPriceValidate_invalidRecommended_shouldReturnError(t *testing.T) {
	price := validPrice()
	price.Recommended = -5
	err := price.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid price recommended" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPriceValidate_zeroCreatedAt_shouldReturnError(t *testing.T) {
	price := validPrice()
	price.CreatedAt = time.Time{}
	err := price.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid price created at" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPriceValidate_zeroUpdatedAt_shouldReturnError(t *testing.T) {
	price := validPrice()
	price.UpdatedAt = time.Time{}
	err := price.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid price updated at" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPriceDelete_shouldSetDeletedAt(t *testing.T) {
	price := validPrice()
	if price.DeletedAt != nil {
		t.Fatal("expected DeletedAt to be nil before Delete")
	}

	before := time.Now()
	price.Delete()
	after := time.Now()

	if price.DeletedAt == nil {
		t.Fatal("expected DeletedAt to be set after Delete")
	}
	if price.DeletedAt.Before(before) || price.DeletedAt.After(after) {
		t.Error("DeletedAt should be between before and after timestamps")
	}
}

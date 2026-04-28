package domain

import (
	"testing"
	"time"
)

func TestNewCost_shouldSetFields(t *testing.T) {
	cost := NewCost("id1", "c1", "b1", "Shipping", 15.50)

	if cost.Id != "id1" {
		t.Errorf("expected Id id1, got %s", cost.Id)
	}
	if cost.CompanyId != "c1" {
		t.Errorf("expected CompanyId c1, got %s", cost.CompanyId)
	}
	if cost.BranchId != "b1" {
		t.Errorf("expected BranchId b1, got %s", cost.BranchId)
	}
	if cost.Name != "Shipping" {
		t.Errorf("expected Name Shipping, got %s", cost.Name)
	}
	if cost.Value != 15.50 {
		t.Errorf("expected Value 15.50, got %f", cost.Value)
	}
	if cost.DeletedAt != nil {
		t.Error("expected DeletedAt to be nil")
	}
}

func validCost() *Cost {
	return NewCost("id1", "c1", "b1", "Shipping", 15.50)
}

func TestCostValidate_valid_shouldReturnNil(t *testing.T) {
	cost := validCost()
	if err := cost.Validate(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestCostValidate_emptyCompanyId_shouldReturnError(t *testing.T) {
	cost := validCost()
	cost.CompanyId = ""
	err := cost.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid cost company id" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCostValidate_emptyBranchId_shouldReturnError(t *testing.T) {
	cost := validCost()
	cost.BranchId = ""
	err := cost.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid cost branch id" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCostValidate_emptyName_shouldReturnError(t *testing.T) {
	cost := validCost()
	cost.Name = ""
	err := cost.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid cost name" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCostValidate_negativeValue_shouldReturnError(t *testing.T) {
	cost := validCost()
	cost.Value = -1.0
	err := cost.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid cost value" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCostDelete_shouldSetDeletedAt(t *testing.T) {
	cost := validCost()
	if cost.DeletedAt != nil {
		t.Fatal("expected DeletedAt to be nil before Delete")
	}

	before := time.Now()
	cost.Delete()
	after := time.Now()

	if cost.DeletedAt == nil {
		t.Fatal("expected DeletedAt to be set after Delete")
	}
	if cost.DeletedAt.Before(before) || cost.DeletedAt.After(after) {
		t.Error("DeletedAt should be between before and after timestamps")
	}
}

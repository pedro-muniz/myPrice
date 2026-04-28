package domain

import (
	"testing"
	"time"
)

func TestNewCompany_shouldSetFields(t *testing.T) {
	company := NewCompany("c1", "Acme Corp")

	if company.Id != "c1" {
		t.Errorf("expected Id c1, got %s", company.Id)
	}
	if company.Name != "Acme Corp" {
		t.Errorf("expected Name Acme Corp, got %s", company.Name)
	}
	if company.DeletedAt != nil {
		t.Error("expected DeletedAt to be nil")
	}
}

func TestCompanyValidate_valid_shouldReturnNil(t *testing.T) {
	company := NewCompany("c1", "Acme Corp")
	if err := company.Validate(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestCompanyValidate_emptyName_shouldReturnError(t *testing.T) {
	company := NewCompany("c1", "")
	err := company.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid company name" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCompanyDelete_shouldSetDeletedAt(t *testing.T) {
	company := NewCompany("c1", "Acme Corp")
	if company.DeletedAt != nil {
		t.Fatal("expected DeletedAt to be nil before Delete")
	}

	deletedAt := time.Now()
	company.Delete(deletedAt)

	if company.DeletedAt == nil {
		t.Fatal("expected DeletedAt to be set after Delete")
	}
	if !company.DeletedAt.Equal(deletedAt) {
		t.Errorf("expected DeletedAt %v, got %v", deletedAt, *company.DeletedAt)
	}
}

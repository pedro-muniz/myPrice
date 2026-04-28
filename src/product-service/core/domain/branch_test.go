package domain

import (
	"testing"
	"time"
)

func TestNewBranch_shouldSetFields(t *testing.T) {
	branch := NewBranch("b1", "c1", "Main Branch")

	if branch.Id != "b1" {
		t.Errorf("expected Id b1, got %s", branch.Id)
	}
	if branch.CompanyId != "c1" {
		t.Errorf("expected CompanyId c1, got %s", branch.CompanyId)
	}
	if branch.Name != "Main Branch" {
		t.Errorf("expected Name Main Branch, got %s", branch.Name)
	}
	if branch.DeletedAt != nil {
		t.Error("expected DeletedAt to be nil")
	}
}

func TestBranchValidate_valid_shouldReturnNil(t *testing.T) {
	branch := NewBranch("b1", "c1", "Main Branch")
	if err := branch.Validate(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestBranchValidate_emptyCompanyId_shouldReturnError(t *testing.T) {
	branch := NewBranch("b1", "", "Main Branch")
	err := branch.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid company id" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestBranchValidate_emptyName_shouldReturnError(t *testing.T) {
	branch := NewBranch("b1", "c1", "")
	err := branch.Validate()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != "invalid branch name" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestBranchDelete_shouldSetDeletedAt(t *testing.T) {
	branch := NewBranch("b1", "c1", "Main Branch")
	if branch.DeletedAt != nil {
		t.Fatal("expected DeletedAt to be nil before Delete")
	}

	deletedAt := time.Now()
	branch.Delete(deletedAt)

	if branch.DeletedAt == nil {
		t.Fatal("expected DeletedAt to be set after Delete")
	}
	if !branch.DeletedAt.Equal(deletedAt) {
		t.Errorf("expected DeletedAt %v, got %v", deletedAt, *branch.DeletedAt)
	}
}

package branch

import (
	"errors"
	"testing"
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	port "github.com/pedro-muniz/myPrice/src/product-service/core/port/usecase/branch"
)

// --- Create ---

func TestBranchManagementCreate_invalidInput_shouldReturnError(t *testing.T) {
	input := &port.CreateInput{CompanyId: "c1", Name: ""} // Invalid: empty name
	uc := &BranchManagement{}

	_, err := uc.Create(input)

	if err == nil {
		t.Error("expected error for invalid input")
	}
}

func TestBranchManagementCreate_repoError_shouldReturnError(t *testing.T) {
	input := &port.CreateInput{CompanyId: "c1", Name: "Main"}
	repoErr := errors.New("db error")
	repository := &TestBranchRepository{
		FakeSave: func(branch *domain.Branch) (<-chan *domain.Branch, <-chan error) {
			ec := make(chan error, 1)
			ec <- repoErr
			return nil, ec
		},
	}
	uc := &BranchManagement{BranchRepository: repository}

	_, err := uc.Create(input)

	if err == nil {
		t.Error("expected error from repository")
	}
}

func TestBranchManagementCreate_success_shouldReturnId(t *testing.T) {
	input := &port.CreateInput{CompanyId: "c1", Name: "Main"}
	savedBranch := &domain.Branch{Id: "branch-123"}
	repository := &TestBranchRepository{
		FakeSave: func(branch *domain.Branch) (<-chan *domain.Branch, <-chan error) {
			bc := make(chan *domain.Branch, 1)
			bc <- savedBranch
			return bc, nil
		},
	}
	uc := &BranchManagement{BranchRepository: repository}

	output, err := uc.Create(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if output.Id != "branch-123" {
		t.Errorf("expected id branch-123, got %s", output.Id)
	}
}

// --- Get ---

func TestBranchManagementGet_nilInput_shouldReturnError(t *testing.T) {
	uc := &BranchManagement{}
	_, err := uc.Get(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestBranchManagementGet_emptyId_shouldReturnError(t *testing.T) {
	input := &port.GetInput{CompanyId: "c1", Id: ""}
	uc := &BranchManagement{}

	_, err := uc.Get(input)

	if err == nil {
		t.Error("expected error for empty id")
	}
}

func TestBranchManagementGet_success_shouldReturnBranch(t *testing.T) {
	existingBranch := &domain.Branch{Id: "b1", CompanyId: "c1", Name: "Main"}
	repository := &TestBranchRepository{
		FakeFindById: func(companyId string, id string) (<-chan *domain.Branch, <-chan error) {
			bc := make(chan *domain.Branch, 1)
			bc <- existingBranch
			return bc, nil
		},
	}
	uc := &BranchManagement{BranchRepository: repository}

	output, err := uc.Get(&port.GetInput{CompanyId: "c1", Id: "b1"})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if output.Branch.Id != "b1" {
		t.Errorf("expected id b1, got %s", output.Branch.Id)
	}
}

// --- List ---

func TestBranchManagementList_nilInput_shouldReturnError(t *testing.T) {
	uc := &BranchManagement{}
	_, err := uc.List(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestBranchManagementList_success_shouldReturnBranches(t *testing.T) {
	branches := []*domain.Branch{
		{Id: "b1", CompanyId: "c1", Name: "Main"},
		{Id: "b2", CompanyId: "c1", Name: "Secondary"},
	}
	repository := &TestBranchRepository{
		FakeFindAll: func(companyId string) (<-chan []*domain.Branch, <-chan error) {
			bc := make(chan []*domain.Branch, 1)
			bc <- branches
			return bc, nil
		},
	}
	uc := &BranchManagement{BranchRepository: repository}

	output, err := uc.List(&port.ListInput{CompanyId: "c1"})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(output.Branches) != 2 {
		t.Errorf("expected 2 branches, got %d", len(output.Branches))
	}
}

// --- Update ---

func TestBranchManagementUpdate_nilInput_shouldReturnError(t *testing.T) {
	uc := &BranchManagement{}
	_, err := uc.Update(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestBranchManagementUpdate_emptyId_shouldReturnError(t *testing.T) {
	input := &port.UpdateInput{CompanyId: "c1", Id: "", Name: "Main"}
	uc := &BranchManagement{}

	_, err := uc.Update(input)

	if err == nil {
		t.Error("expected error for empty id")
	}
}

func TestBranchManagementUpdate_invalidInput_shouldReturnError(t *testing.T) {
	input := &port.UpdateInput{CompanyId: "c1", Id: "b1", Name: ""} // Invalid: empty name
	existingBranch := &domain.Branch{Id: "b1", CompanyId: "c1", Name: "Old Name"}
	repository := &TestBranchRepository{
		FakeFindById: func(companyId string, id string) (<-chan *domain.Branch, <-chan error) {
			bc := make(chan *domain.Branch, 1)
			bc <- existingBranch
			return bc, nil
		},
	}
	uc := &BranchManagement{BranchRepository: repository}

	_, err := uc.Update(input)

	if err == nil {
		t.Error("expected error for invalid input")
	}
}

func TestBranchManagementUpdate_success_shouldReturnSuccess(t *testing.T) {
	input := &port.UpdateInput{CompanyId: "c1", Id: "b1", Name: "New Name"}
	existingBranch := &domain.Branch{Id: "b1", CompanyId: "c1", Name: "Old Name"}
	repository := &TestBranchRepository{
		FakeFindById: func(companyId string, id string) (<-chan *domain.Branch, <-chan error) {
			bc := make(chan *domain.Branch, 1)
			bc <- existingBranch
			return bc, nil
		},
		FakeUpdate: func(branch *domain.Branch) <-chan error {
			ec := make(chan error, 1)
			ec <- nil
			return ec
		},
	}
	uc := &BranchManagement{BranchRepository: repository}

	output, err := uc.Update(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !output.Success {
		t.Error("expected success true")
	}
}

// --- Delete ---

func TestBranchManagementDelete_nilInput_shouldReturnError(t *testing.T) {
	uc := &BranchManagement{}
	_, err := uc.Delete(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestBranchManagementDelete_emptyId_shouldReturnError(t *testing.T) {
	input := &port.DeleteInput{CompanyId: "c1", Id: ""}
	uc := &BranchManagement{}

	_, err := uc.Delete(input)

	if err == nil {
		t.Error("expected error for empty id")
	}
}

func TestBranchManagementDelete_repoError_shouldReturnError(t *testing.T) {
	input := &port.DeleteInput{CompanyId: "c1", Id: "b1"}
	repository := &TestBranchRepository{
		FakeDelete: func(companyId string, id string, deletedAt time.Time) <-chan error {
			ec := make(chan error, 1)
			ec <- errors.New("db error")
			return ec
		},
	}
	uc := &BranchManagement{BranchRepository: repository}

	_, err := uc.Delete(input)

	if err == nil {
		t.Error("expected error from repository")
	}
}

func TestBranchManagementDelete_success_shouldReturnSuccess(t *testing.T) {
	input := &port.DeleteInput{CompanyId: "c1", Id: "b1"}
	repository := &TestBranchRepository{
		FakeDelete: func(companyId string, id string, deletedAt time.Time) <-chan error {
			ec := make(chan error, 1)
			ec <- nil
			return ec
		},
	}
	uc := &BranchManagement{BranchRepository: repository}

	output, err := uc.Delete(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !output.Success {
		t.Error("expected success true")
	}
}

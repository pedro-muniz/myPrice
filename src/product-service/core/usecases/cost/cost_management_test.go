package cost

import (
	"errors"
	"testing"
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	port "github.com/pedro-muniz/myPrice/src/product-service/core/port/usecase/cost"
)

type TestCostRepository struct {
	FakeSave     func(cost *domain.Cost) (<-chan *domain.Cost, <-chan error)
	FakeFindById func(companyId string, branchId string, id string) (<-chan *domain.Cost, <-chan error)
	FakeFindAll  func(companyId string, branchId string) (<-chan []*domain.Cost, <-chan error)
	FakeUpdate   func(cost *domain.Cost) <-chan error
	FakeDelete   func(companyId string, branchId string, id string, deletedAt time.Time) <-chan error
}

func (r *TestCostRepository) Save(cost *domain.Cost) (<-chan *domain.Cost, <-chan error) {
	return r.FakeSave(cost)
}
func (r *TestCostRepository) FindById(companyId string, branchId string, id string) (<-chan *domain.Cost, <-chan error) {
	return r.FakeFindById(companyId, branchId, id)
}
func (r *TestCostRepository) FindAll(companyId string, branchId string) (<-chan []*domain.Cost, <-chan error) {
	return r.FakeFindAll(companyId, branchId)
}
func (r *TestCostRepository) Update(cost *domain.Cost) <-chan error {
	return r.FakeUpdate(cost)
}
func (r *TestCostRepository) Delete(companyId string, branchId string, id string, deletedAt time.Time) <-chan error {
	return r.FakeDelete(companyId, branchId, id, deletedAt)
}

// --- Create ---

func TestCostManagementCreate_nilInput_shouldReturnError(t *testing.T) {
	uc := &CostManagement{}
	_, err := uc.Create(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestCostManagementCreate_invalidInput_shouldReturnError(t *testing.T) {
	input := &port.CreateInput{CompanyId: "c1", BranchId: "b1", Name: "", Value: 10.0} // Invalid: empty name
	uc := &CostManagement{}

	_, err := uc.Create(input)

	if err == nil {
		t.Error("expected error for invalid input")
	}
}

func TestCostManagementCreate_repoError_shouldReturnError(t *testing.T) {
	input := &port.CreateInput{CompanyId: "c1", BranchId: "b1", Name: "Cost", Value: 10.0}
	repoErr := errors.New("db error")
	repository := &TestCostRepository{
		FakeSave: func(cost *domain.Cost) (<-chan *domain.Cost, <-chan error) {
			ec := make(chan error, 1)
			ec <- repoErr
			return nil, ec
		},
	}
	uc := &CostManagement{CostRepository: repository}

	_, err := uc.Create(input)

	if err == nil {
		t.Error("expected error from repository")
	}
}

func TestCostManagementCreate_success_shouldReturnId(t *testing.T) {
	input := &port.CreateInput{CompanyId: "c1", BranchId: "b1", Name: "Cost", Value: 10.0}
	savedCost := &domain.Cost{Id: "cost-123"}
	repository := &TestCostRepository{
		FakeSave: func(cost *domain.Cost) (<-chan *domain.Cost, <-chan error) {
			cc := make(chan *domain.Cost, 1)
			cc <- savedCost
			return cc, nil
		},
	}
	uc := &CostManagement{CostRepository: repository}

	output, err := uc.Create(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if output.Id != "cost-123" {
		t.Errorf("expected id cost-123, got %s", output.Id)
	}
}

// --- Get ---

func TestCostManagementGet_nilInput_shouldReturnError(t *testing.T) {
	uc := &CostManagement{}
	_, err := uc.Get(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestCostManagementGet_emptyId_shouldReturnError(t *testing.T) {
	input := &port.GetInput{CompanyId: "c1", BranchId: "b1", Id: ""}
	uc := &CostManagement{}

	_, err := uc.Get(input)

	if err == nil {
		t.Error("expected error for empty id")
	}
}

func TestCostManagementGet_success_shouldReturnCost(t *testing.T) {
	existingCost := &domain.Cost{Id: "id1"}
	repository := &TestCostRepository{
		FakeFindById: func(companyId string, branchId string, id string) (<-chan *domain.Cost, <-chan error) {
			cc := make(chan *domain.Cost, 1)
			cc <- existingCost
			return cc, nil
		},
	}
	uc := &CostManagement{CostRepository: repository}

	output, err := uc.Get(&port.GetInput{CompanyId: "c1", BranchId: "b1", Id: "id1"})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if output.Cost.Id != "id1" {
		t.Errorf("expected id id1, got %s", output.Cost.Id)
	}
}

// --- List ---

func TestCostManagementList_nilInput_shouldReturnError(t *testing.T) {
	uc := &CostManagement{}
	_, err := uc.List(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestCostManagementList_success_shouldReturnCosts(t *testing.T) {
	costs := []*domain.Cost{{Id: "id1"}, {Id: "id2"}}
	repository := &TestCostRepository{
		FakeFindAll: func(companyId string, branchId string) (<-chan []*domain.Cost, <-chan error) {
			cc := make(chan []*domain.Cost, 1)
			cc <- costs
			return cc, nil
		},
	}
	uc := &CostManagement{CostRepository: repository}

	output, err := uc.List(&port.ListInput{CompanyId: "c1", BranchId: "b1"})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(output.Costs) != 2 {
		t.Errorf("expected 2 costs, got %d", len(output.Costs))
	}
}

// --- Update ---

func TestCostManagementUpdate_nilInput_shouldReturnError(t *testing.T) {
	uc := &CostManagement{}
	_, err := uc.Update(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestCostManagementUpdate_emptyId_shouldReturnError(t *testing.T) {
	input := &port.UpdateInput{CompanyId: "c1", BranchId: "b1", Id: ""}
	uc := &CostManagement{}

	_, err := uc.Update(input)

	if err == nil {
		t.Error("expected error for empty id")
	}
}

func TestCostManagementUpdate_invalidInput_shouldReturnError(t *testing.T) {
	input := &port.UpdateInput{CompanyId: "c1", BranchId: "b1", Id: "id1", Name: ""}
	existingCost := &domain.Cost{Id: "id1", CompanyId: "c1", BranchId: "b1", Name: "Old"}
	repository := &TestCostRepository{
		FakeFindById: func(companyId string, branchId string, id string) (<-chan *domain.Cost, <-chan error) {
			cc := make(chan *domain.Cost, 1)
			cc <- existingCost
			return cc, nil
		},
	}
	uc := &CostManagement{CostRepository: repository}

	_, err := uc.Update(input)

	if err == nil {
		t.Error("expected error for invalid input")
	}
}

func TestCostManagementUpdate_success_shouldReturnSuccess(t *testing.T) {
	input := &port.UpdateInput{CompanyId: "c1", BranchId: "b1", Id: "id1", Name: "New", Value: 20.0}
	existingCost := &domain.Cost{Id: "id1", CompanyId: "c1", BranchId: "b1", Name: "Old"}
	repository := &TestCostRepository{
		FakeFindById: func(companyId string, branchId string, id string) (<-chan *domain.Cost, <-chan error) {
			cc := make(chan *domain.Cost, 1)
			cc <- existingCost
			return cc, nil
		},
		FakeUpdate: func(cost *domain.Cost) <-chan error {
			ec := make(chan error, 1)
			ec <- nil
			return ec
		},
	}
	uc := &CostManagement{CostRepository: repository}

	output, err := uc.Update(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !output.Success {
		t.Error("expected success true")
	}
}

// --- Delete ---

func TestCostManagementDelete_nilInput_shouldReturnError(t *testing.T) {
	uc := &CostManagement{}
	_, err := uc.Delete(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestCostManagementDelete_emptyId_shouldReturnError(t *testing.T) {
	input := &port.DeleteInput{CompanyId: "c1", BranchId: "b1", Id: ""}
	uc := &CostManagement{}

	_, err := uc.Delete(input)

	if err == nil {
		t.Error("expected error for empty id")
	}
}

func TestCostManagementDelete_repoError_shouldReturnError(t *testing.T) {
	input := &port.DeleteInput{CompanyId: "c1", BranchId: "b1", Id: "id1"}
	repository := &TestCostRepository{
		FakeDelete: func(companyId string, branchId string, id string, deletedAt time.Time) <-chan error {
			ec := make(chan error, 1)
			ec <- errors.New("db error")
			return ec
		},
	}
	uc := &CostManagement{CostRepository: repository}

	_, err := uc.Delete(input)

	if err == nil {
		t.Error("expected error from repository")
	}
}

func TestCostManagementDelete_success_shouldReturnSuccess(t *testing.T) {
	input := &port.DeleteInput{CompanyId: "c1", BranchId: "b1", Id: "id1"}
	repository := &TestCostRepository{
		FakeDelete: func(companyId string, branchId string, id string, deletedAt time.Time) <-chan error {
			ec := make(chan error, 1)
			ec <- nil
			return ec
		},
	}
	uc := &CostManagement{CostRepository: repository}

	output, err := uc.Delete(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !output.Success {
		t.Error("expected success true")
	}
}

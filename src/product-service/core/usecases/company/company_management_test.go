package company

import (
	"errors"
	"testing"
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	port "github.com/pedro-muniz/myPrice/src/product-service/core/port/usecase/company"
)

// --- Create ---

func TestCompanyManagementCreate_invalidInput_shouldReturnError(t *testing.T) {
	input := &port.CreateInput{Name: ""} // Invalid: empty name
	uc := &CompanyManagement{}

	_, err := uc.Create(input)

	if err == nil {
		t.Error("expected error for invalid input")
	}
}

func TestCompanyManagementCreate_repoError_shouldReturnError(t *testing.T) {
	input := &port.CreateInput{Name: "Acme Corp"}
	repoErr := errors.New("db error")
	repository := &TestCompanyRepository{
		FakeSave: func(company *domain.Company) (<-chan *domain.Company, <-chan error) {
			ec := make(chan error, 1)
			ec <- repoErr
			return nil, ec
		},
	}
	uc := &CompanyManagement{CompanyRepository: repository}

	_, err := uc.Create(input)

	if err == nil {
		t.Error("expected error from repository")
	}
}

func TestCompanyManagementCreate_success_shouldReturnId(t *testing.T) {
	input := &port.CreateInput{Name: "Acme Corp"}
	savedCompany := &domain.Company{Id: "company-123"}
	repository := &TestCompanyRepository{
		FakeSave: func(company *domain.Company) (<-chan *domain.Company, <-chan error) {
			cc := make(chan *domain.Company, 1)
			cc <- savedCompany
			return cc, nil
		},
	}
	uc := &CompanyManagement{CompanyRepository: repository}

	output, err := uc.Create(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if output.Id != "company-123" {
		t.Errorf("expected id company-123, got %s", output.Id)
	}
}

// --- Get ---

func TestCompanyManagementGet_nilInput_shouldReturnError(t *testing.T) {
	uc := &CompanyManagement{}
	_, err := uc.Get(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestCompanyManagementGet_emptyId_shouldReturnError(t *testing.T) {
	input := &port.GetInput{Id: ""}
	uc := &CompanyManagement{}

	_, err := uc.Get(input)

	if err == nil {
		t.Error("expected error for empty id")
	}
}

func TestCompanyManagementGet_success_shouldReturnCompany(t *testing.T) {
	existingCompany := &domain.Company{Id: "c1", Name: "Acme Corp"}
	repository := &TestCompanyRepository{
		FakeFindById: func(id string) (<-chan *domain.Company, <-chan error) {
			cc := make(chan *domain.Company, 1)
			cc <- existingCompany
			return cc, nil
		},
	}
	uc := &CompanyManagement{CompanyRepository: repository}

	output, err := uc.Get(&port.GetInput{Id: "c1"})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if output.Company.Id != "c1" {
		t.Errorf("expected id c1, got %s", output.Company.Id)
	}
}

// --- List ---

func TestCompanyManagementList_nilInput_shouldReturnError(t *testing.T) {
	uc := &CompanyManagement{}
	_, err := uc.List(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestCompanyManagementList_success_shouldReturnCompanies(t *testing.T) {
	companies := []*domain.Company{
		{Id: "c1", Name: "Acme"},
		{Id: "c2", Name: "Globex"},
	}
	repository := &TestCompanyRepository{
		FakeFindAll: func() (<-chan []*domain.Company, <-chan error) {
			cc := make(chan []*domain.Company, 1)
			cc <- companies
			return cc, nil
		},
	}
	uc := &CompanyManagement{CompanyRepository: repository}

	output, err := uc.List(&port.ListInput{})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(output.Companies) != 2 {
		t.Errorf("expected 2 companies, got %d", len(output.Companies))
	}
}

// --- Update ---

func TestCompanyManagementUpdate_nilInput_shouldReturnError(t *testing.T) {
	uc := &CompanyManagement{}
	_, err := uc.Update(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestCompanyManagementUpdate_emptyId_shouldReturnError(t *testing.T) {
	input := &port.UpdateInput{Id: "", Name: "Acme"}
	uc := &CompanyManagement{}

	_, err := uc.Update(input)

	if err == nil {
		t.Error("expected error for empty id")
	}
}

func TestCompanyManagementUpdate_invalidInput_shouldReturnError(t *testing.T) {
	input := &port.UpdateInput{Id: "c1", Name: ""} // Invalid: empty name
	existingCompany := &domain.Company{Id: "c1", Name: "Old Name"}
	repository := &TestCompanyRepository{
		FakeFindById: func(id string) (<-chan *domain.Company, <-chan error) {
			cc := make(chan *domain.Company, 1)
			cc <- existingCompany
			return cc, nil
		},
	}
	uc := &CompanyManagement{CompanyRepository: repository}

	_, err := uc.Update(input)

	if err == nil {
		t.Error("expected error for invalid input")
	}
}

func TestCompanyManagementUpdate_success_shouldReturnSuccess(t *testing.T) {
	input := &port.UpdateInput{Id: "c1", Name: "New Name"}
	existingCompany := &domain.Company{Id: "c1", Name: "Old Name"}
	repository := &TestCompanyRepository{
		FakeFindById: func(id string) (<-chan *domain.Company, <-chan error) {
			cc := make(chan *domain.Company, 1)
			cc <- existingCompany
			return cc, nil
		},
		FakeUpdate: func(company *domain.Company) <-chan error {
			ec := make(chan error, 1)
			ec <- nil
			return ec
		},
	}
	uc := &CompanyManagement{CompanyRepository: repository}

	output, err := uc.Update(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !output.Success {
		t.Error("expected success true")
	}
}

// --- Delete ---

func TestCompanyManagementDelete_nilInput_shouldReturnError(t *testing.T) {
	uc := &CompanyManagement{}
	_, err := uc.Delete(nil)
	if err == nil {
		t.Error("expected error for nil input")
	}
}

func TestCompanyManagementDelete_emptyId_shouldReturnError(t *testing.T) {
	input := &port.DeleteInput{Id: ""}
	uc := &CompanyManagement{}

	_, err := uc.Delete(input)

	if err == nil {
		t.Error("expected error for empty id")
	}
}

func TestCompanyManagementDelete_repoError_shouldReturnError(t *testing.T) {
	input := &port.DeleteInput{Id: "c1"}
	repository := &TestCompanyRepository{
		FakeDelete: func(id string, deletedAt time.Time) <-chan error {
			ec := make(chan error, 1)
			ec <- errors.New("db error")
			return ec
		},
	}
	uc := &CompanyManagement{CompanyRepository: repository}

	_, err := uc.Delete(input)

	if err == nil {
		t.Error("expected error from repository")
	}
}

func TestCompanyManagementDelete_success_shouldReturnSuccess(t *testing.T) {
	input := &port.DeleteInput{Id: "c1"}
	repository := &TestCompanyRepository{
		FakeDelete: func(id string, deletedAt time.Time) <-chan error {
			ec := make(chan error, 1)
			ec <- nil
			return ec
		},
	}
	uc := &CompanyManagement{CompanyRepository: repository}

	output, err := uc.Delete(input)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !output.Success {
		t.Error("expected success true")
	}
}

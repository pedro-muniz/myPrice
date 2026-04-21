package company

import (
	"sync"
	"time"

	companyErrors "github.com/pedro-muniz/myPrice/src/product-service/core/customerror/company"
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	repository "github.com/pedro-muniz/myPrice/src/product-service/core/port/repository"
	port "github.com/pedro-muniz/myPrice/src/product-service/core/port/usecase/company"
)

type CompanyManagement struct {
	CompanyRepository repository.ICompanyRepository
}

var companyManagementOnce sync.Once
var companyManagementInstance *CompanyManagement

func GetCompanyManagementInstance(companyRepository repository.ICompanyRepository) *CompanyManagement {
	companyManagementOnce.Do(func() {
		companyManagementInstance = &CompanyManagement{
			CompanyRepository: companyRepository,
		}
	})

	return companyManagementInstance
}

func (this *CompanyManagement) Create(input *port.CreateInput) (*port.CreateOutput, error) {
	company := domain.NewCompany("", input.Name)

	if err := company.Validate(); err != nil {
		return nil, companyErrors.InvalidCompanyData(err.Error())
	}

	companyChan, errChan := this.CompanyRepository.Save(company)

	var savedCompany *domain.Company
	var err error

	select {
	case savedCompany = <-companyChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, companyErrors.ErrorSavingCompanyDatabaseRecord(err.Error())
	}

	if savedCompany == nil {
		return nil, companyErrors.ErrorSavingCompanyDatabaseRecord("saved company is nil")
	}

	return &port.CreateOutput{
		Id: savedCompany.Id,
	}, nil
}

func (this *CompanyManagement) Get(input *port.GetInput) (*port.GetOutput, error) {
	if input == nil || input.Id == "" {
		return nil, companyErrors.InvalidCompanyData("invalid id")
	}

	companyChan, errChan := this.CompanyRepository.FindById(input.Id)

	var company *domain.Company
	var err error

	select {
	case company = <-companyChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, companyErrors.ErrorGettingCompanyDatabaseRecord(err.Error())
	}

	if company == nil {
		return nil, companyErrors.CompanyNotFound(input.Id)
	}

	return &port.GetOutput{
		Company: company,
	}, nil
}

func (this *CompanyManagement) List(input *port.ListInput) (*port.ListOutput, error) {
	if input == nil {
		return nil, companyErrors.InvalidCompanyData("nil input")
	}

	companiesChan, errChan := this.CompanyRepository.FindAll()

	var companies []*domain.Company
	var err error

	select {
	case companies = <-companiesChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, companyErrors.ErrorGettingCompanyDatabaseRecord(err.Error())
	}

	return &port.ListOutput{
		Companies: companies,
	}, nil
}

func (this *CompanyManagement) Update(input *port.UpdateInput) (*port.UpdateOutput, error) {
	if input == nil || input.Id == "" {
		return nil, companyErrors.InvalidCompanyData("invalid id")
	}

	companyChan, findErrChan := this.CompanyRepository.FindById(input.Id)
	var existingCompany *domain.Company
	var findErr error

	select {
	case existingCompany = <-companyChan:
	case findErr = <-findErrChan:
	}

	if findErr != nil {
		return nil, companyErrors.ErrorGettingCompanyDatabaseRecord(findErr.Error())
	}
	if existingCompany == nil {
		return nil, companyErrors.CompanyNotFound(input.Id)
	}

	existingCompany.Name = input.Name

	if err := existingCompany.Validate(); err != nil {
		return nil, companyErrors.InvalidCompanyData(err.Error())
	}

	errChan := this.CompanyRepository.Update(existingCompany)

	var err error = <-errChan
	if err != nil {
		return nil, companyErrors.ErrorSavingCompanyDatabaseRecord(err.Error())
	}

	return &port.UpdateOutput{
		Success: true,
	}, nil
}

func (this *CompanyManagement) Delete(input *port.DeleteInput) (*port.DeleteOutput, error) {
	if input == nil || input.Id == "" {
		return nil, companyErrors.InvalidCompanyData("invalid id")
	}

	company := &domain.Company{
		Id: input.Id,
	}
	company.Delete(time.Now())

	errChan := this.CompanyRepository.Delete(company.Id, *company.DeletedAt)

	var err error = <-errChan
	if err != nil {
		return nil, companyErrors.ErrorSavingCompanyDatabaseRecord(err.Error())
	}

	return &port.DeleteOutput{
		Success: true,
	}, nil
}

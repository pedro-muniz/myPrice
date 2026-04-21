package branch

import (
	"sync"
	"time"

	branchErrors "github.com/pedro-muniz/myPrice/src/product-service/core/customerror/branch"
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	repository "github.com/pedro-muniz/myPrice/src/product-service/core/port/repository"
	port "github.com/pedro-muniz/myPrice/src/product-service/core/port/usecase/branch"
)

type BranchManagement struct {
	BranchRepository repository.IBranchRepository
}

var branchManagementOnce sync.Once
var branchManagementInstance *BranchManagement

func GetBranchManagementInstance(branchRepository repository.IBranchRepository) *BranchManagement {
	branchManagementOnce.Do(func() {
		branchManagementInstance = &BranchManagement{
			BranchRepository: branchRepository,
		}
	})

	return branchManagementInstance
}

func (this *BranchManagement) Create(input *port.CreateInput) (*port.CreateOutput, error) {
	branch := domain.NewBranch("", input.CompanyId, input.Name)

	if err := branch.Validate(); err != nil {
		return nil, branchErrors.InvalidBranchData(err.Error())
	}

	branchChan, errChan := this.BranchRepository.Save(branch)

	var savedBranch *domain.Branch
	var err error

	select {
	case savedBranch = <-branchChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, branchErrors.ErrorSavingBranchDatabaseRecord(err.Error())
	}

	if savedBranch == nil {
		return nil, branchErrors.ErrorSavingBranchDatabaseRecord("saved branch is nil")
	}

	return &port.CreateOutput{
		Id: savedBranch.Id,
	}, nil
}

func (this *BranchManagement) Get(input *port.GetInput) (*port.GetOutput, error) {
	if input == nil || input.Id == "" {
		return nil, branchErrors.InvalidBranchData("invalid id")
	}

	branchChan, errChan := this.BranchRepository.FindById(input.CompanyId, input.Id)

	var branch *domain.Branch
	var err error

	select {
	case branch = <-branchChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, branchErrors.ErrorGettingBranchDatabaseRecord(err.Error())
	}

	if branch == nil {
		return nil, branchErrors.BranchNotFound(input.Id)
	}

	return &port.GetOutput{
		Branch: branch,
	}, nil
}

func (this *BranchManagement) List(input *port.ListInput) (*port.ListOutput, error) {
	if input == nil {
		return nil, branchErrors.InvalidBranchData("nil input")
	}

	branchesChan, errChan := this.BranchRepository.FindAll(input.CompanyId)

	var branches []*domain.Branch
	var err error

	select {
	case branches = <-branchesChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, branchErrors.ErrorGettingBranchDatabaseRecord(err.Error())
	}

	return &port.ListOutput{
		Branches: branches,
	}, nil
}

func (this *BranchManagement) Update(input *port.UpdateInput) (*port.UpdateOutput, error) {
	if input == nil || input.Id == "" {
		return nil, branchErrors.InvalidBranchData("invalid id")
	}

	branchChan, findErrChan := this.BranchRepository.FindById(input.CompanyId, input.Id)
	var existingBranch *domain.Branch
	var findErr error

	select {
	case existingBranch = <-branchChan:
	case findErr = <-findErrChan:
	}

	if findErr != nil {
		return nil, branchErrors.ErrorGettingBranchDatabaseRecord(findErr.Error())
	}
	if existingBranch == nil {
		return nil, branchErrors.BranchNotFound(input.Id)
	}

	existingBranch.Name = input.Name

	if err := existingBranch.Validate(); err != nil {
		return nil, branchErrors.InvalidBranchData(err.Error())
	}

	errChan := this.BranchRepository.Update(existingBranch)

	var err error = <-errChan
	if err != nil {
		return nil, branchErrors.ErrorSavingBranchDatabaseRecord(err.Error())
	}

	return &port.UpdateOutput{
		Success: true,
	}, nil
}

func (this *BranchManagement) Delete(input *port.DeleteInput) (*port.DeleteOutput, error) {
	if input == nil || input.Id == "" {
		return nil, branchErrors.InvalidBranchData("invalid id")
	}

	branch := &domain.Branch{
		Id:        input.Id,
		CompanyId: input.CompanyId,
	}
	branch.Delete(time.Now())

	errChan := this.BranchRepository.Delete(branch.CompanyId, branch.Id, *branch.DeletedAt)

	var err error = <-errChan
	if err != nil {
		return nil, branchErrors.ErrorSavingBranchDatabaseRecord(err.Error())
	}

	return &port.DeleteOutput{
		Success: true,
	}, nil
}

package cost

import (
	"sync"
	"time"

	costErrors "github.com/pedro-muniz/myPrice/src/product-service/core/customerror/cost"
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	repository "github.com/pedro-muniz/myPrice/src/product-service/core/port/repository"
	port "github.com/pedro-muniz/myPrice/src/product-service/core/port/usecase/cost"
)

type CostManagement struct {
	CostRepository repository.ICostRepository
}

var costManagementOnce sync.Once
var costManagementInstance *CostManagement

func GetCostManagementInstance(costRepository repository.ICostRepository) *CostManagement {
	costManagementOnce.Do(func() {
		costManagementInstance = &CostManagement{
			CostRepository: costRepository,
		}
	})

	return costManagementInstance
}

func (this *CostManagement) Create(input *port.CreateInput) (*port.CreateOutput, error) {
	if input == nil {
		return nil, costErrors.InvalidCostData("nil input")
	}

	cost := domain.NewCost("", input.CompanyId, input.BranchId, input.Name, input.Value)

	if err := cost.Validate(); err != nil {
		return nil, costErrors.InvalidCostData(err.Error())
	}

	costChan, errChan := this.CostRepository.Save(cost)

	var savedCost *domain.Cost
	var err error

	select {
	case savedCost = <-costChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, costErrors.ErrorSavingCostDatabaseRecord(err.Error())
	}

	if savedCost == nil {
		return nil, costErrors.ErrorSavingCostDatabaseRecord("saved cost is nil")
	}

	return &port.CreateOutput{
		Id: savedCost.Id,
	}, nil
}

func (this *CostManagement) Get(input *port.GetInput) (*port.GetOutput, error) {
	if input == nil || input.Id == "" {
		return nil, costErrors.InvalidCostData("invalid id")
	}

	costChan, errChan := this.CostRepository.FindById(input.CompanyId, input.BranchId, input.Id)

	var cost *domain.Cost
	var err error

	select {
	case cost = <-costChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, costErrors.ErrorGettingCostDatabaseRecord(err.Error())
	}

	if cost == nil {
		return nil, costErrors.CostNotFound(input.Id)
	}

	return &port.GetOutput{
		Cost: cost,
	}, nil
}

func (this *CostManagement) List(input *port.ListInput) (*port.ListOutput, error) {
	if input == nil {
		return nil, costErrors.InvalidCostData("nil input")
	}

	costsChan, errChan := this.CostRepository.FindAll(input.CompanyId, input.BranchId)

	var costs []*domain.Cost
	var err error

	select {
	case costs = <-costsChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, costErrors.ErrorGettingCostDatabaseRecord(err.Error())
	}

	return &port.ListOutput{
		Costs: costs,
	}, nil
}

func (this *CostManagement) Update(input *port.UpdateInput) (*port.UpdateOutput, error) {
	if input == nil || input.Id == "" {
		return nil, costErrors.InvalidCostData("invalid id")
	}

	costChan, findErrChan := this.CostRepository.FindById(input.CompanyId, input.BranchId, input.Id)
	var existingCost *domain.Cost
	var findErr error

	select {
	case existingCost = <-costChan:
	case findErr = <-findErrChan:
	}

	if findErr != nil {
		return nil, costErrors.ErrorGettingCostDatabaseRecord(findErr.Error())
	}
	if existingCost == nil {
		return nil, costErrors.CostNotFound(input.Id)
	}

	existingCost.Name = input.Name
	existingCost.Value = input.Value
	existingCost.UpdatedAt = time.Now()

	if err := existingCost.Validate(); err != nil {
		return nil, costErrors.InvalidCostData(err.Error())
	}

	errChan := this.CostRepository.Update(existingCost)

	var err error = <-errChan
	if err != nil {
		return nil, costErrors.ErrorSavingCostDatabaseRecord(err.Error())
	}

	return &port.UpdateOutput{
		Success: true,
	}, nil
}

func (this *CostManagement) Delete(input *port.DeleteInput) (*port.DeleteOutput, error) {
	if input == nil || input.Id == "" {
		return nil, costErrors.InvalidCostData("invalid id")
	}

	cost := &domain.Cost{
		Id:        input.Id,
		CompanyId: input.CompanyId,
		BranchId:  input.BranchId,
	}
	cost.Delete()

	errChan := this.CostRepository.Delete(cost.CompanyId, cost.BranchId, cost.Id, *cost.DeletedAt)

	var err error = <-errChan
	if err != nil {
		return nil, costErrors.ErrorSavingCostDatabaseRecord(err.Error())
	}

	return &port.DeleteOutput{
		Success: true,
	}, nil
}

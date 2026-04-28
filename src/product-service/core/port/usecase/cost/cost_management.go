package cost

import (
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type CreateInput struct {
	CompanyId string
	BranchId  string
	Name      string
	Value     float64
}

type CreateOutput struct {
	Id string
}

type GetInput struct {
	CompanyId string
	BranchId  string
	Id        string
}

type GetOutput struct {
	Cost *domain.Cost
}

type ListInput struct {
	CompanyId string
	BranchId  string
}

type ListOutput struct {
	Costs []*domain.Cost
}

type UpdateInput struct {
	CompanyId string
	BranchId  string
	Id        string
	Name      string
	Value     float64
}

type UpdateOutput struct {
	Success bool
}

type DeleteInput struct {
	CompanyId string
	BranchId  string
	Id        string
}

type DeleteOutput struct {
	Success bool
}

type ICostManagement interface {
	Create(input *CreateInput) (*CreateOutput, error)
	Get(input *GetInput) (*GetOutput, error)
	List(input *ListInput) (*ListOutput, error)
	Update(input *UpdateInput) (*UpdateOutput, error)
	Delete(input *DeleteInput) (*DeleteOutput, error)
}

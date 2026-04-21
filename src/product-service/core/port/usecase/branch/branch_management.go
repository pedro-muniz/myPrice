package branch

import (
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type CreateInput struct {
	CompanyId string
	Name      string
}

type CreateOutput struct {
	Id string
}

type GetInput struct {
	CompanyId string
	Id        string
}

type GetOutput struct {
	Branch *domain.Branch
}

type ListInput struct {
	CompanyId string
}

type ListOutput struct {
	Branches []*domain.Branch
}

type UpdateInput struct {
	CompanyId string
	Id        string
	Name      string
}

type UpdateOutput struct {
	Success bool
}

type DeleteInput struct {
	CompanyId string
	Id        string
}

type DeleteOutput struct {
	Success bool
}

type IBranchManagement interface {
	Create(input *CreateInput) (*CreateOutput, error)
	Get(input *GetInput) (*GetOutput, error)
	List(input *ListInput) (*ListOutput, error)
	Update(input *UpdateInput) (*UpdateOutput, error)
	Delete(input *DeleteInput) (*DeleteOutput, error)
}

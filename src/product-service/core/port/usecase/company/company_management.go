package company

import (
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type CreateInput struct {
	Name string
}

type CreateOutput struct {
	Id string
}

type GetInput struct {
	Id string
}

type GetOutput struct {
	Company *domain.Company
}

type ListInput struct{}

type ListOutput struct {
	Companies []*domain.Company
}

type UpdateInput struct {
	Id   string
	Name string
}

type UpdateOutput struct {
	Success bool
}

type DeleteInput struct {
	Id string
}

type DeleteOutput struct {
	Success bool
}

type ICompanyManagement interface {
	Create(input *CreateInput) (*CreateOutput, error)
	Get(input *GetInput) (*GetOutput, error)
	List(input *ListInput) (*ListOutput, error)
	Update(input *UpdateInput) (*UpdateOutput, error)
	Delete(input *DeleteInput) (*DeleteOutput, error)
}

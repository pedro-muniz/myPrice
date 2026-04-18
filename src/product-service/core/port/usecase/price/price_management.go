package price

import (
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type CreateInput struct {
	Gross       float64
	Net         float64
	Selling     float64
	Recommended float64
}

type CreateOutput struct {
	Id string
}

type GetInput struct {
	Id string
}

type GetOutput struct {
	Price *domain.Price
}

type ListInput struct {
}

type ListOutput struct {
	Prices []*domain.Price
}

type UpdateInput struct {
	Id          string
	Gross       float64
	Net         float64
	Selling     float64
	Recommended float64
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

type IPriceManagement interface {
	Create(input *CreateInput) (*CreateOutput, error)
	Get(input *GetInput) (*GetOutput, error)
	List(input *ListInput) (*ListOutput, error)
	Update(input *UpdateInput) (*UpdateOutput, error)
	Delete(input *DeleteInput) (*DeleteOutput, error)
}

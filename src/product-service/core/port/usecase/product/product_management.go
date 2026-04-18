package product

import (
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type CreateInput struct {
	BarCode       string
	Name          string
	Description   string
	Discount      float64
	ProfitPercent float64
	Origin        int
	Gross         float64
	Net           float64
	Selling       float64
	Recommended   float64
}

type CreateOutput struct {
	Id string
}

type GetInput struct {
	Id string
}

type GetOutput struct {
	Product *domain.Product
}

type ListInput struct {
}

type ListOutput struct {
	Products []*domain.Product
}

type UpdateInput struct {
	Id            string
	BarCode       string
	Name          string
	Description   string
	Discount      float64
	ProfitPercent float64
	Origin        int
	Gross         float64
	Net           float64
	Selling       float64
	Recommended   float64
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

type IProductManagement interface {
	Create(input *CreateInput) (*CreateOutput, error)
	Get(input *GetInput) (*GetOutput, error)
	List(input *ListInput) (*ListOutput, error)
	Update(input *UpdateInput) (*UpdateOutput, error)
	Delete(input *DeleteInput) (*DeleteOutput, error)
}

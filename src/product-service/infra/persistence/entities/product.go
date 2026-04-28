package entities

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type ProductEntity struct {
	Id            string
	CompanyId     string
	BranchId      string
	BarCode       string
	Name          string
	Description   string
	Discount      float64
	PriceId       string
	ProfitPercent float64
	Origin        int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

func (e *ProductEntity) ToDomain(price *domain.Price) *domain.Product {
	return &domain.Product{
		Id:            e.Id,
		CompanyId:     e.CompanyId,
		BranchId:      e.BranchId,
		BarCode:       e.BarCode,
		Name:          e.Name,
		Description:   e.Description,
		Discount:      e.Discount,
		Price:         *price,
		ProfitPercent: e.ProfitPercent,
		Origin:        domain.Origin(e.Origin),
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		DeletedAt:     e.DeletedAt,
	}
}

func (e *ProductEntity) FromDomain(d *domain.Product) {
	e.Id = d.Id
	e.CompanyId = d.CompanyId
	e.BranchId = d.BranchId
	e.BarCode = d.BarCode
	e.Name = d.Name
	e.Description = d.Description
	e.Discount = d.Discount
	e.PriceId = d.Price.Id
	e.ProfitPercent = d.ProfitPercent
	e.Origin = int(d.Origin)
	e.CreatedAt = d.CreatedAt
	e.UpdatedAt = d.UpdatedAt
	e.DeletedAt = d.DeletedAt
}

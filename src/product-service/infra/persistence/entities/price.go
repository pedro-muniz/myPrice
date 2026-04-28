package entities

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type PriceEntity struct {
	Id          string
	CompanyId   string
	BranchId    string
	Gross       float64
	Net         float64
	Selling     float64
	Recommended float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (e *PriceEntity) ToDomain() *domain.Price {
	return &domain.Price{
		Id:          e.Id,
		CompanyId:   e.CompanyId,
		BranchId:    e.BranchId,
		Gross:       e.Gross,
		Net:         e.Net,
		Selling:     e.Selling,
		Recommended: e.Recommended,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		DeletedAt:   e.DeletedAt,
	}
}

func (e *PriceEntity) FromDomain(d *domain.Price) {
	e.Id = d.Id
	e.CompanyId = d.CompanyId
	e.BranchId = d.BranchId
	e.Gross = d.Gross
	e.Net = d.Net
	e.Selling = d.Selling
	e.Recommended = d.Recommended
	e.CreatedAt = d.CreatedAt
	e.UpdatedAt = d.UpdatedAt
	e.DeletedAt = d.DeletedAt
}

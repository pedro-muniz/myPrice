package entities

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type CostEntity struct {
	Id        string
	CompanyId string
	BranchId  string
	Name      string
	Value     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (e *CostEntity) ToDomain() *domain.Cost {
	return &domain.Cost{
		Id:        e.Id,
		CompanyId: e.CompanyId,
		BranchId:  e.BranchId,
		Name:      e.Name,
		Value:     e.Value,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: e.DeletedAt,
	}
}

func (e *CostEntity) FromDomain(d *domain.Cost) {
	e.Id = d.Id
	e.CompanyId = d.CompanyId
	e.BranchId = d.BranchId
	e.Name = d.Name
	e.Value = d.Value
	e.CreatedAt = d.CreatedAt
	e.UpdatedAt = d.UpdatedAt
	e.DeletedAt = d.DeletedAt
}

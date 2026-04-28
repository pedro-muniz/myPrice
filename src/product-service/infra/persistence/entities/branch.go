package entities

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type BranchEntity struct {
	Id        string
	CompanyId string
	Name      string
	DeletedAt *time.Time
}

func (e *BranchEntity) ToDomain() *domain.Branch {
	return &domain.Branch{
		Id:        e.Id,
		CompanyId: e.CompanyId,
		Name:      e.Name,
		DeletedAt: e.DeletedAt,
	}
}

func (e *BranchEntity) FromDomain(d *domain.Branch) {
	e.Id = d.Id
	e.CompanyId = d.CompanyId
	e.Name = d.Name
	e.DeletedAt = d.DeletedAt
}

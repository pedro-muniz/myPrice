package entities

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type CompanyEntity struct {
	Id        string
	Name      string
	DeletedAt *time.Time
}

func (e *CompanyEntity) ToDomain() *domain.Company {
	return &domain.Company{
		Id:        e.Id,
		Name:      e.Name,
		DeletedAt: e.DeletedAt,
	}
}

func (e *CompanyEntity) FromDomain(d *domain.Company) {
	e.Id = d.Id
	e.Name = d.Name
	e.DeletedAt = d.DeletedAt
}

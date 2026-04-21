package repository

import (
	time "time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type ICompanyRepository interface {
	Save(company *domain.Company) (<-chan *domain.Company, <-chan error)
	FindById(id string) (<-chan *domain.Company, <-chan error)
	FindAll() (<-chan []*domain.Company, <-chan error)
	Update(company *domain.Company) <-chan error
	Delete(id string, deletedAt time.Time) <-chan error
}

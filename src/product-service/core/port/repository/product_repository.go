package repository

import (
	time "time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type IProductRepository interface {
	Save(product *domain.Product) (<-chan *domain.Product, <-chan error)
	FindById(companyId string, branchId string, id string) (<-chan *domain.Product, <-chan error)
	FindAll(companyId string, branchId string) (<-chan []*domain.Product, <-chan error)
	Update(product *domain.Product) <-chan error
	Delete(companyId string, branchId string, id string, deletedAt time.Time) <-chan error
}

package repository

import (
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type IProductRepository interface {
	Save(product *domain.Product) (<-chan *domain.Product, <-chan error)
	FindById(id string) (<-chan *domain.Product, <-chan error)
	FindAll() (<-chan []*domain.Product, <-chan error)
	Update(product *domain.Product) <-chan error
	Delete(id string) <-chan error
}

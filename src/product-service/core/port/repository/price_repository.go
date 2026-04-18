package repository

import (
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type IPriceRepository interface {
	Save(price *domain.Price) (<-chan *domain.Price, <-chan error)
	FindById(id string) (<-chan *domain.Price, <-chan error)
	FindAll() (<-chan []*domain.Price, <-chan error)
	Update(price *domain.Price) <-chan error
	Delete(id string) <-chan error
}

package repository

import (
	time "time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type IPriceRepository interface {
	Save(price *domain.Price) (<-chan *domain.Price, <-chan error)
	FindById(companyId string, branchId string, id string) (<-chan *domain.Price, <-chan error)
	FindAll(companyId string, branchId string) (<-chan []*domain.Price, <-chan error)
	Update(price *domain.Price) <-chan error
	Delete(companyId string, branchId string, id string, deletedAt time.Time) <-chan error
}

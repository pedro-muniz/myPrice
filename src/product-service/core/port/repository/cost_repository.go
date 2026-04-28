package repository

import (
	time "time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type ICostRepository interface {
	Save(cost *domain.Cost) (<-chan *domain.Cost, <-chan error)
	FindById(companyId string, branchId string, id string) (<-chan *domain.Cost, <-chan error)
	FindAll(companyId string, branchId string) (<-chan []*domain.Cost, <-chan error)
	Update(cost *domain.Cost) <-chan error
	Delete(companyId string, branchId string, id string, deletedAt time.Time) <-chan error
}

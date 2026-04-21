package repository

import (
	time "time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type IBranchRepository interface {
	Save(branch *domain.Branch) (<-chan *domain.Branch, <-chan error)
	FindById(companyId string, id string) (<-chan *domain.Branch, <-chan error)
	FindAll(companyId string) (<-chan []*domain.Branch, <-chan error)
	Update(branch *domain.Branch) <-chan error
	Delete(companyId string, id string, deletedAt time.Time) <-chan error
}

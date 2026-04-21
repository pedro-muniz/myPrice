package product

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	_ "github.com/pedro-muniz/myPrice/src/product-service/core/port/repository"
)

type TestProductRepository struct {
	FakeSave     func(product *domain.Product) (<-chan *domain.Product, <-chan error)
	FakeFindById func(companyId string, branchId string, id string) (<-chan *domain.Product, <-chan error)
	FakeFindAll  func(companyId string, branchId string) (<-chan []*domain.Product, <-chan error)
	FakeUpdate   func(product *domain.Product) <-chan error
	FakeDelete   func(companyId string, branchId string, id string, deletedAt time.Time) <-chan error
}

func (m *TestProductRepository) Save(product *domain.Product) (<-chan *domain.Product, <-chan error) {
	return m.FakeSave(product)
}

func (m *TestProductRepository) FindById(companyId string, branchId string, id string) (<-chan *domain.Product, <-chan error) {
	return m.FakeFindById(companyId, branchId, id)
}

func (m *TestProductRepository) FindAll(companyId string, branchId string) (<-chan []*domain.Product, <-chan error) {
	return m.FakeFindAll(companyId, branchId)
}

func (m *TestProductRepository) Update(product *domain.Product) <-chan error {
	return m.FakeUpdate(product)
}

func (m *TestProductRepository) Delete(companyId string, branchId string, id string, deletedAt time.Time) <-chan error {
	return m.FakeDelete(companyId, branchId, id, deletedAt)
}

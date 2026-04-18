package product

import (
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	_ "github.com/pedro-muniz/myPrice/src/product-service/core/port/repository"
)

type TestProductRepository struct {
	FakeSave     func(product *domain.Product) (<-chan *domain.Product, <-chan error)
	FakeFindById func(id string) (<-chan *domain.Product, <-chan error)
	FakeFindAll  func() (<-chan []*domain.Product, <-chan error)
	FakeUpdate   func(product *domain.Product) <-chan error
	FakeDelete   func(id string) <-chan error
}

func (m *TestProductRepository) Save(product *domain.Product) (<-chan *domain.Product, <-chan error) {
	return m.FakeSave(product)
}

func (m *TestProductRepository) FindById(id string) (<-chan *domain.Product, <-chan error) {
	return m.FakeFindById(id)
}

func (m *TestProductRepository) FindAll() (<-chan []*domain.Product, <-chan error) {
	return m.FakeFindAll()
}

func (m *TestProductRepository) Update(product *domain.Product) <-chan error {
	return m.FakeUpdate(product)
}

func (m *TestProductRepository) Delete(id string) <-chan error {
	return m.FakeDelete(id)
}

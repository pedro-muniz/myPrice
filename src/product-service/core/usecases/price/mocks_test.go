package price

import (
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type TestPriceRepository struct {
	FakeSave     func(price *domain.Price) (<-chan *domain.Price, <-chan error)
	FakeFindById func(id string) (<-chan *domain.Price, <-chan error)
	FakeFindAll  func() (<-chan []*domain.Price, <-chan error)
	FakeUpdate   func(price *domain.Price) <-chan error
	FakeDelete   func(id string) <-chan error
}

func (m *TestPriceRepository) Save(price *domain.Price) (<-chan *domain.Price, <-chan error) {
	return m.FakeSave(price)
}

func (m *TestPriceRepository) FindById(id string) (<-chan *domain.Price, <-chan error) {
	return m.FakeFindById(id)
}

func (m *TestPriceRepository) FindAll() (<-chan []*domain.Price, <-chan error) {
	return m.FakeFindAll()
}

func (m *TestPriceRepository) Update(price *domain.Price) <-chan error {
	return m.FakeUpdate(price)
}

func (m *TestPriceRepository) Delete(id string) <-chan error {
	return m.FakeDelete(id)
}

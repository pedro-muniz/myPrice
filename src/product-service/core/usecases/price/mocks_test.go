package price

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type TestPriceRepository struct {
	FakeSave     func(price *domain.Price) (<-chan *domain.Price, <-chan error)
	FakeFindById func(companyId string, branchId string, id string) (<-chan *domain.Price, <-chan error)
	FakeFindAll  func(companyId string, branchId string) (<-chan []*domain.Price, <-chan error)
	FakeUpdate   func(price *domain.Price) <-chan error
	FakeDelete   func(companyId string, branchId string, id string, deletedAt time.Time) <-chan error
}

func (m *TestPriceRepository) Save(price *domain.Price) (<-chan *domain.Price, <-chan error) {
	return m.FakeSave(price)
}

func (m *TestPriceRepository) FindById(companyId string, branchId string, id string) (<-chan *domain.Price, <-chan error) {
	return m.FakeFindById(companyId, branchId, id)
}

func (m *TestPriceRepository) FindAll(companyId string, branchId string) (<-chan []*domain.Price, <-chan error) {
	return m.FakeFindAll(companyId, branchId)
}

func (m *TestPriceRepository) Update(price *domain.Price) <-chan error {
	return m.FakeUpdate(price)
}

func (m *TestPriceRepository) Delete(companyId string, branchId string, id string, deletedAt time.Time) <-chan error {
	return m.FakeDelete(companyId, branchId, id, deletedAt)
}

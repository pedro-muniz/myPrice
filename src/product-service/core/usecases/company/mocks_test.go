package company

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type TestCompanyRepository struct {
	FakeSave     func(company *domain.Company) (<-chan *domain.Company, <-chan error)
	FakeFindById func(id string) (<-chan *domain.Company, <-chan error)
	FakeFindAll  func() (<-chan []*domain.Company, <-chan error)
	FakeUpdate   func(company *domain.Company) <-chan error
	FakeDelete   func(id string, deletedAt time.Time) <-chan error
}

func (m *TestCompanyRepository) Save(company *domain.Company) (<-chan *domain.Company, <-chan error) {
	return m.FakeSave(company)
}

func (m *TestCompanyRepository) FindById(id string) (<-chan *domain.Company, <-chan error) {
	return m.FakeFindById(id)
}

func (m *TestCompanyRepository) FindAll() (<-chan []*domain.Company, <-chan error) {
	return m.FakeFindAll()
}

func (m *TestCompanyRepository) Update(company *domain.Company) <-chan error {
	return m.FakeUpdate(company)
}

func (m *TestCompanyRepository) Delete(id string, deletedAt time.Time) <-chan error {
	return m.FakeDelete(id, deletedAt)
}

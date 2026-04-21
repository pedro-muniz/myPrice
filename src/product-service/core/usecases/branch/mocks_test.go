package branch

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
)

type TestBranchRepository struct {
	FakeSave     func(branch *domain.Branch) (<-chan *domain.Branch, <-chan error)
	FakeFindById func(companyId string, id string) (<-chan *domain.Branch, <-chan error)
	FakeFindAll  func(companyId string) (<-chan []*domain.Branch, <-chan error)
	FakeUpdate   func(branch *domain.Branch) <-chan error
	FakeDelete   func(companyId string, id string, deletedAt time.Time) <-chan error
}

func (m *TestBranchRepository) Save(branch *domain.Branch) (<-chan *domain.Branch, <-chan error) {
	return m.FakeSave(branch)
}

func (m *TestBranchRepository) FindById(companyId string, id string) (<-chan *domain.Branch, <-chan error) {
	return m.FakeFindById(companyId, id)
}

func (m *TestBranchRepository) FindAll(companyId string) (<-chan []*domain.Branch, <-chan error) {
	return m.FakeFindAll(companyId)
}

func (m *TestBranchRepository) Update(branch *domain.Branch) <-chan error {
	return m.FakeUpdate(branch)
}

func (m *TestBranchRepository) Delete(companyId string, id string, deletedAt time.Time) <-chan error {
	return m.FakeDelete(companyId, id, deletedAt)
}

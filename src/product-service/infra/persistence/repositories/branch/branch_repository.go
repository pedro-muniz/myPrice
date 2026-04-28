package branch

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	entities "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/entities"
	repositories "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories"
)

type BranchQueries interface {
	Insert() string
	FindById() string
	FindAll() string
	Update() string
	Delete() string
}

type BranchRepository struct {
	dao     repositories.IDAO
	queries BranchQueries
}

func NewBranchRepository(dao repositories.IDAO, queries BranchQueries) *BranchRepository {
	return &BranchRepository{dao: dao, queries: queries}
}

func (r *BranchRepository) Save(branch *domain.Branch) (<-chan *domain.Branch, <-chan error) {
	branchChan := make(chan *domain.Branch, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(branchChan)
		defer close(errChan)

		var dbBranch entities.BranchEntity
		dbBranch.FromDomain(branch)
		_, err := r.dao.Write(r.queries.Insert(), dbBranch.CompanyId, dbBranch.Name)
		if err != nil {
			errChan <- err
			return
		}

		branchChan <- branch
	}()

	return branchChan, errChan
}

func (r *BranchRepository) FindById(companyId string, id string) (<-chan *domain.Branch, <-chan error) {
	branchChan := make(chan *domain.Branch, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(branchChan)
		defer close(errChan)

		rows, err := r.dao.Read(r.queries.FindById(), companyId, id)
		if err != nil {
			errChan <- err
			return
		}
		defer rows.Close()

		if rows.Next() {
			var dbBranch entities.BranchEntity
			if err := rows.Scan(&dbBranch.Id, &dbBranch.CompanyId, &dbBranch.Name, &dbBranch.DeletedAt); err != nil {
				errChan <- err
				return
			}
			branchChan <- dbBranch.ToDomain()
		} else {
			branchChan <- nil
		}
	}()

	return branchChan, errChan
}

func (r *BranchRepository) FindAll(companyId string) (<-chan []*domain.Branch, <-chan error) {
	branchesChan := make(chan []*domain.Branch, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(branchesChan)
		defer close(errChan)

		rows, err := r.dao.Read(r.queries.FindAll(), companyId)
		if err != nil {
			errChan <- err
			return
		}
		defer rows.Close()

		var branches []*domain.Branch
		for rows.Next() {
			var dbBranch entities.BranchEntity
			if err := rows.Scan(&dbBranch.Id, &dbBranch.CompanyId, &dbBranch.Name, &dbBranch.DeletedAt); err != nil {
				errChan <- err
				return
			}
			branches = append(branches, dbBranch.ToDomain())
		}
		branchesChan <- branches
	}()

	return branchesChan, errChan
}

func (r *BranchRepository) Update(branch *domain.Branch) <-chan error {
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)

		var dbBranch entities.BranchEntity
		dbBranch.FromDomain(branch)
		_, err := r.dao.Write(r.queries.Update(), dbBranch.Name, dbBranch.CompanyId, dbBranch.Id)
		if err != nil {
			errChan <- err
		}
	}()

	return errChan
}

func (r *BranchRepository) Delete(companyId string, id string, deletedAt time.Time) <-chan error {
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)

		_, err := r.dao.Write(r.queries.Delete(), deletedAt, companyId, id)
		if err != nil {
			errChan <- err
		}
	}()

	return errChan
}

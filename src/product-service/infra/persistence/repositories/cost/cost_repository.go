package cost

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	entities "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/entities"
	repositories "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories"
)

type CostQueries interface {
	Insert() string
	FindById() string
	FindAll() string
	Update() string
	Delete() string
}

type CostRepository struct {
	dao     repositories.IDAO
	queries CostQueries
}

func NewCostRepository(dao repositories.IDAO, queries CostQueries) *CostRepository {
	return &CostRepository{dao: dao, queries: queries}
}

func (r *CostRepository) Save(cost *domain.Cost) (<-chan *domain.Cost, <-chan error) {
	costChan := make(chan *domain.Cost, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(costChan)
		defer close(errChan)

		var dbCost entities.CostEntity
		dbCost.FromDomain(cost)
		_, err := r.dao.Write(r.queries.Insert(), dbCost.CompanyId, dbCost.BranchId, dbCost.Name, dbCost.Value)
		if err != nil {
			errChan <- err
			return
		}

		costChan <- cost
	}()

	return costChan, errChan
}

func (r *CostRepository) FindById(companyId string, branchId string, id string) (<-chan *domain.Cost, <-chan error) {
	costChan := make(chan *domain.Cost, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(costChan)
		defer close(errChan)

		rows, err := r.dao.Read(r.queries.FindById(), companyId, branchId, id)
		if err != nil {
			errChan <- err
			return
		}
		defer rows.Close()

		if rows.Next() {
			var dbCost entities.CostEntity
			if err := rows.Scan(&dbCost.Id, &dbCost.CompanyId, &dbCost.BranchId, &dbCost.Name, &dbCost.Value,
				&dbCost.CreatedAt, &dbCost.UpdatedAt, &dbCost.DeletedAt); err != nil {
				errChan <- err
				return
			}
			costChan <- dbCost.ToDomain()
		} else {
			costChan <- nil
		}
	}()

	return costChan, errChan
}

func (r *CostRepository) FindAll(companyId string, branchId string) (<-chan []*domain.Cost, <-chan error) {
	costsChan := make(chan []*domain.Cost, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(costsChan)
		defer close(errChan)

		rows, err := r.dao.Read(r.queries.FindAll(), companyId, branchId)
		if err != nil {
			errChan <- err
			return
		}
		defer rows.Close()

		var costs []*domain.Cost
		for rows.Next() {
			var dbCost entities.CostEntity
			if err := rows.Scan(&dbCost.Id, &dbCost.CompanyId, &dbCost.BranchId, &dbCost.Name, &dbCost.Value,
				&dbCost.CreatedAt, &dbCost.UpdatedAt, &dbCost.DeletedAt); err != nil {
				errChan <- err
				return
			}
			costs = append(costs, dbCost.ToDomain())
		}
		costsChan <- costs
	}()

	return costsChan, errChan
}

func (r *CostRepository) Update(cost *domain.Cost) <-chan error {
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)

		var dbCost entities.CostEntity
		dbCost.FromDomain(cost)
		_, err := r.dao.Write(r.queries.Update(), dbCost.Name, dbCost.Value, dbCost.UpdatedAt,
			dbCost.CompanyId, dbCost.BranchId, dbCost.Id)
		if err != nil {
			errChan <- err
		}
	}()

	return errChan
}

func (r *CostRepository) Delete(companyId string, branchId string, id string, deletedAt time.Time) <-chan error {
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)

		_, err := r.dao.Write(r.queries.Delete(), deletedAt, companyId, branchId, id)
		if err != nil {
			errChan <- err
		}
	}()

	return errChan
}

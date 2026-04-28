package price

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	entities "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/entities"
	repositories "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories"
)

type PriceQueries interface {
	Insert() string
	FindById() string
	FindAll() string
	Update() string
	Delete() string
}

type PriceRepository struct {
	dao     repositories.IDAO
	queries PriceQueries
}

func NewPriceRepository(dao repositories.IDAO, queries PriceQueries) *PriceRepository {
	return &PriceRepository{dao: dao, queries: queries}
}

func (r *PriceRepository) Save(price *domain.Price) (<-chan *domain.Price, <-chan error) {
	priceChan := make(chan *domain.Price, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(priceChan)
		defer close(errChan)

		var dbPrice entities.PriceEntity
		dbPrice.FromDomain(price)
		_, err := r.dao.Write(r.queries.Insert(), dbPrice.CompanyId, dbPrice.BranchId,
			dbPrice.Gross, dbPrice.Net, dbPrice.Selling, dbPrice.Recommended, dbPrice.CreatedAt, dbPrice.UpdatedAt)
		if err != nil {
			errChan <- err
			return
		}

		priceChan <- price
	}()

	return priceChan, errChan
}

func (r *PriceRepository) FindById(companyId string, branchId string, id string) (<-chan *domain.Price, <-chan error) {
	priceChan := make(chan *domain.Price, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(priceChan)
		defer close(errChan)

		rows, err := r.dao.Read(r.queries.FindById(), companyId, branchId, id)
		if err != nil {
			errChan <- err
			return
		}
		defer rows.Close()

		if rows.Next() {
			var dbPrice entities.PriceEntity
			if err := rows.Scan(&dbPrice.Id, &dbPrice.CompanyId, &dbPrice.BranchId, &dbPrice.Gross,
				&dbPrice.Net, &dbPrice.Selling, &dbPrice.Recommended, &dbPrice.CreatedAt,
				&dbPrice.UpdatedAt, &dbPrice.DeletedAt); err != nil {
				errChan <- err
				return
			}
			priceChan <- dbPrice.ToDomain()
		} else {
			priceChan <- nil
		}
	}()

	return priceChan, errChan
}

func (r *PriceRepository) FindAll(companyId string, branchId string) (<-chan []*domain.Price, <-chan error) {
	pricesChan := make(chan []*domain.Price, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(pricesChan)
		defer close(errChan)

		rows, err := r.dao.Read(r.queries.FindAll(), companyId, branchId)
		if err != nil {
			errChan <- err
			return
		}
		defer rows.Close()

		var prices []*domain.Price
		for rows.Next() {
			var dbPrice entities.PriceEntity
			if err := rows.Scan(&dbPrice.Id, &dbPrice.CompanyId, &dbPrice.BranchId, &dbPrice.Gross,
				&dbPrice.Net, &dbPrice.Selling, &dbPrice.Recommended, &dbPrice.CreatedAt,
				&dbPrice.UpdatedAt, &dbPrice.DeletedAt); err != nil {
				errChan <- err
				return
			}
			prices = append(prices, dbPrice.ToDomain())
		}
		pricesChan <- prices
	}()

	return pricesChan, errChan
}

func (r *PriceRepository) Update(price *domain.Price) <-chan error {
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)

		var dbPrice entities.PriceEntity
		dbPrice.FromDomain(price)
		_, err := r.dao.Write(r.queries.Update(), dbPrice.Gross, dbPrice.Net, dbPrice.Selling,
			dbPrice.Recommended, dbPrice.UpdatedAt, dbPrice.CompanyId, dbPrice.BranchId, dbPrice.Id)
		if err != nil {
			errChan <- err
		}
	}()

	return errChan
}

func (r *PriceRepository) Delete(companyId string, branchId string, id string, deletedAt time.Time) <-chan error {
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

package company

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	entities "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/entities"
	repositories "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories"
)

type CompanyQueries interface {
	Insert() string
	FindById() string
	FindAll() string
	Update() string
	Delete() string
}

type CompanyRepository struct {
	dao     repositories.IDAO
	queries CompanyQueries
}

func NewCompanyRepository(dao repositories.IDAO, queries CompanyQueries) *CompanyRepository {
	return &CompanyRepository{dao: dao, queries: queries}
}

func (r *CompanyRepository) Save(company *domain.Company) (<-chan *domain.Company, <-chan error) {
	companyChan := make(chan *domain.Company, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(companyChan)
		defer close(errChan)

		var dbCompany entities.CompanyEntity
		dbCompany.FromDomain(company)
		_, err := r.dao.Write(r.queries.Insert(), dbCompany.Name)
		if err != nil {
			errChan <- err
			return
		}

		companyChan <- company
	}()

	return companyChan, errChan
}

func (r *CompanyRepository) FindById(id string) (<-chan *domain.Company, <-chan error) {
	companyChan := make(chan *domain.Company, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(companyChan)
		defer close(errChan)

		rows, err := r.dao.Read(r.queries.FindById(), id)
		if err != nil {
			errChan <- err
			return
		}
		defer rows.Close()

		if rows.Next() {
			var dbCompany entities.CompanyEntity
			if err := rows.Scan(&dbCompany.Id, &dbCompany.Name, &dbCompany.DeletedAt); err != nil {
				errChan <- err
				return
			}
			companyChan <- dbCompany.ToDomain()
		} else {
			companyChan <- nil
		}
	}()

	return companyChan, errChan
}

func (r *CompanyRepository) FindAll() (<-chan []*domain.Company, <-chan error) {
	companiesChan := make(chan []*domain.Company, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(companiesChan)
		defer close(errChan)

		rows, err := r.dao.Read(r.queries.FindAll())
		if err != nil {
			errChan <- err
			return
		}
		defer rows.Close()

		var companies []*domain.Company
		for rows.Next() {
			var dbCompany entities.CompanyEntity
			if err := rows.Scan(&dbCompany.Id, &dbCompany.Name, &dbCompany.DeletedAt); err != nil {
				errChan <- err
				return
			}
			companies = append(companies, dbCompany.ToDomain())
		}
		companiesChan <- companies
	}()

	return companiesChan, errChan
}

func (r *CompanyRepository) Update(company *domain.Company) <-chan error {
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)

		var dbCompany entities.CompanyEntity
		dbCompany.FromDomain(company)
		_, err := r.dao.Write(r.queries.Update(), dbCompany.Name, dbCompany.Id)
		if err != nil {
			errChan <- err
		}
	}()

	return errChan
}

func (r *CompanyRepository) Delete(id string, deletedAt time.Time) <-chan error {
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)

		_, err := r.dao.Write(r.queries.Delete(), deletedAt, id)
		if err != nil {
			errChan <- err
		}
	}()

	return errChan
}

package product

import (
	"time"

	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	entities "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/entities"
	repositories "github.com/pedro-muniz/myPrice/src/product-service/infra/persistence/repositories"
)

type ProductQueries interface {
	Insert() string
	FindById() string
	FindAll() string
	Update() string
	Delete() string
}

type ProductRepository struct {
	dao     repositories.IDAO
	queries ProductQueries
}

func NewProductRepository(dao repositories.IDAO, queries ProductQueries) *ProductRepository {
	return &ProductRepository{dao: dao, queries: queries}
}

func (r *ProductRepository) Save(product *domain.Product) (<-chan *domain.Product, <-chan error) {
	productChan := make(chan *domain.Product, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(productChan)
		defer close(errChan)

		var dbProduct entities.ProductEntity
		dbProduct.FromDomain(product)
		_, err := r.dao.Write(r.queries.Insert(), dbProduct.CompanyId, dbProduct.BranchId, dbProduct.BarCode,
			dbProduct.Name, dbProduct.Description, dbProduct.Discount, dbProduct.PriceId, dbProduct.ProfitPercent,
			dbProduct.Origin, dbProduct.CreatedAt, dbProduct.UpdatedAt)
		if err != nil {
			errChan <- err
			return
		}

		productChan <- product
	}()

	return productChan, errChan
}

func (r *ProductRepository) FindById(companyId string, branchId string, id string) (<-chan *domain.Product, <-chan error) {
	productChan := make(chan *domain.Product, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(productChan)
		defer close(errChan)

		rows, err := r.dao.Read(r.queries.FindById(), companyId, branchId, id)
		if err != nil {
			errChan <- err
			return
		}
		defer rows.Close()

		if rows.Next() {
			var dbProduct entities.ProductEntity
			var dbPrice entities.PriceEntity
			if err := rows.Scan(
				&dbProduct.Id, &dbProduct.CompanyId, &dbProduct.BranchId, &dbProduct.BarCode,
				&dbProduct.Name, &dbProduct.Description, &dbProduct.Discount, &dbProduct.ProfitPercent,
				&dbProduct.Origin, &dbProduct.CreatedAt, &dbProduct.UpdatedAt, &dbProduct.DeletedAt,
				&dbPrice.Id, &dbPrice.Gross, &dbPrice.Net, &dbPrice.Selling, &dbPrice.Recommended,
				&dbPrice.CreatedAt, &dbPrice.UpdatedAt, &dbPrice.DeletedAt,
			); err != nil {
				errChan <- err
				return
			}
			price := dbPrice.ToDomain()
			price.CompanyId = dbProduct.CompanyId
			price.BranchId = dbProduct.BranchId
			productChan <- dbProduct.ToDomain(price)
		} else {
			productChan <- nil
		}
	}()

	return productChan, errChan
}

func (r *ProductRepository) FindAll(companyId string, branchId string) (<-chan []*domain.Product, <-chan error) {
	productsChan := make(chan []*domain.Product, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(productsChan)
		defer close(errChan)

		rows, err := r.dao.Read(r.queries.FindAll(), companyId, branchId)
		if err != nil {
			errChan <- err
			return
		}
		defer rows.Close()

		var products []*domain.Product
		for rows.Next() {
			var dbProduct entities.ProductEntity
			var dbPrice entities.PriceEntity
			if err := rows.Scan(
				&dbProduct.Id, &dbProduct.CompanyId, &dbProduct.BranchId, &dbProduct.BarCode,
				&dbProduct.Name, &dbProduct.Description, &dbProduct.Discount, &dbProduct.ProfitPercent,
				&dbProduct.Origin, &dbProduct.CreatedAt, &dbProduct.UpdatedAt, &dbProduct.DeletedAt,
				&dbPrice.Id, &dbPrice.Gross, &dbPrice.Net, &dbPrice.Selling, &dbPrice.Recommended,
				&dbPrice.CreatedAt, &dbPrice.UpdatedAt, &dbPrice.DeletedAt,
			); err != nil {
				errChan <- err
				return
			}
			price := dbPrice.ToDomain()
			price.CompanyId = dbProduct.CompanyId
			price.BranchId = dbProduct.BranchId
			products = append(products, dbProduct.ToDomain(price))
		}
		productsChan <- products
	}()

	return productsChan, errChan
}

func (r *ProductRepository) Update(product *domain.Product) <-chan error {
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)

		var dbProduct entities.ProductEntity
		dbProduct.FromDomain(product)
		_, err := r.dao.Write(r.queries.Update(), dbProduct.BarCode, dbProduct.Name,
			dbProduct.Description, dbProduct.Discount, dbProduct.ProfitPercent, dbProduct.Origin,
			dbProduct.UpdatedAt, dbProduct.CompanyId, dbProduct.BranchId, dbProduct.Id)
		if err != nil {
			errChan <- err
			return
		}
	}()

	return errChan
}

func (r *ProductRepository) Delete(companyId string, branchId string, id string, deletedAt time.Time) <-chan error {
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

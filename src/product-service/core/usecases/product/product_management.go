package product

import (
	"sync"
	"time"

	productErrors "github.com/pedro-muniz/myPrice/src/product-service/core/customerror/product"
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	repository "github.com/pedro-muniz/myPrice/src/product-service/core/port/repository"
	port "github.com/pedro-muniz/myPrice/src/product-service/core/port/usecase/product"
)

type ProductManagement struct {
	ProductRepository repository.IProductRepository
}

var productManagementOnce sync.Once
var productManagementInstance *ProductManagement

func GetProductManagementInstance(productRepository repository.IProductRepository) *ProductManagement {
	productManagementOnce.Do(func() {
		productManagementInstance = &ProductManagement{
			ProductRepository: productRepository,
		}
	})

	return productManagementInstance
}

func (this *ProductManagement) Create(input *port.CreateInput) (*port.CreateOutput, error) {
	if input == nil {
		return nil, productErrors.InvalidProductData("nil input")
	}

	price := domain.Price{
		CompanyId:   input.CompanyId,
		BranchId:    input.BranchId,
		Gross:       input.Gross,
		Net:         input.Net,
		Selling:     input.Selling,
		Recommended: input.Recommended,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	product := domain.NewProduct(
		input.CompanyId,
		input.BranchId,
		input.BarCode,
		input.Name,
		input.Description,
		input.Discount,
		price,
		input.ProfitPercent,
		domain.Origin(input.Origin),
	)
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	if err := product.Validate(); err != nil {
		return nil, productErrors.InvalidProductData(err.Error())
	}

	productChan, errChan := this.ProductRepository.Save(product)

	var savedProduct *domain.Product
	var err error

	select {
	case savedProduct = <-productChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, productErrors.ErrorSavingProductDatabaseRecord(err.Error())
	}

	if savedProduct == nil {
		return nil, productErrors.ErrorSavingProductDatabaseRecord("saved product is nil")
	}

	return &port.CreateOutput{
		Id: savedProduct.Id,
	}, nil
}

func (this *ProductManagement) Get(input *port.GetInput) (*port.GetOutput, error) {
	if input == nil || input.Id == "" {
		return nil, productErrors.InvalidProductData("invalid id")
	}

	productChan, errChan := this.ProductRepository.FindById(input.CompanyId, input.BranchId, input.Id)

	var product *domain.Product
	var err error

	select {
	case product = <-productChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, productErrors.ErrorGettingProductDatabaseRecord(err.Error())
	}

	if product == nil {
		return nil, productErrors.ProductNotFound(input.Id)
	}

	return &port.GetOutput{
		Product: product,
	}, nil
}

func (this *ProductManagement) List(input *port.ListInput) (*port.ListOutput, error) {
	if input == nil {
		return nil, productErrors.InvalidProductData("nil input")
	}
	productsChan, errChan := this.ProductRepository.FindAll(input.CompanyId, input.BranchId)

	var products []*domain.Product
	var err error

	select {
	case products = <-productsChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, productErrors.ErrorGettingProductDatabaseRecord(err.Error())
	}

	return &port.ListOutput{
		Products: products,
	}, nil
}

func (this *ProductManagement) Update(input *port.UpdateInput) (*port.UpdateOutput, error) {
	if input == nil || input.Id == "" {
		return nil, productErrors.InvalidProductData("invalid id")
	}

	productChan, findErrChan := this.ProductRepository.FindById(input.CompanyId, input.BranchId, input.Id)
	var existingProduct *domain.Product
	var findErr error

	select {
	case existingProduct = <-productChan:
	case findErr = <-findErrChan:
	}

	if findErr != nil {
		return nil, productErrors.ErrorGettingProductDatabaseRecord(findErr.Error())
	}
	if existingProduct == nil {
		return nil, productErrors.ProductNotFound(input.Id)
	}

	existingProduct.BarCode = input.BarCode
	existingProduct.Name = input.Name
	existingProduct.Description = input.Description
	existingProduct.Discount = input.Discount
	existingProduct.ProfitPercent = input.ProfitPercent
	existingProduct.Origin = domain.Origin(input.Origin)
	existingProduct.UpdatedAt = time.Now()

	existingProduct.Price.Gross = input.Gross
	existingProduct.Price.Net = input.Net
	existingProduct.Price.Selling = input.Selling
	existingProduct.Price.Recommended = input.Recommended
	existingProduct.Price.UpdatedAt = time.Now()

	if err := existingProduct.Validate(); err != nil {
		return nil, productErrors.InvalidProductData(err.Error())
	}

	errChan := this.ProductRepository.Update(existingProduct)

	var err error = <-errChan
	if err != nil {
		return nil, productErrors.ErrorSavingProductDatabaseRecord(err.Error())
	}

	return &port.UpdateOutput{
		Success: true,
	}, nil
}

func (this *ProductManagement) Delete(input *port.DeleteInput) (*port.DeleteOutput, error) {
	if input == nil || input.Id == "" {
		return nil, productErrors.InvalidProductData("invalid id")
	}

	product := &domain.Product{
		Id:        input.Id,
		CompanyId: input.CompanyId,
		BranchId:  input.BranchId,
	}
	product.Delete()

	errChan := this.ProductRepository.Delete(product.CompanyId, product.BranchId, product.Id, *product.DeletedAt)

	var err error = <-errChan
	if err != nil {
		return nil, productErrors.ErrorSavingProductDatabaseRecord(err.Error())
	}

	return &port.DeleteOutput{
		Success: true,
	}, nil
}

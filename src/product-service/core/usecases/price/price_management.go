package price

import (
	"sync"
	"time"

	priceErrors "github.com/pedro-muniz/myPrice/src/product-service/core/customerror/price"
	domain "github.com/pedro-muniz/myPrice/src/product-service/core/domain"
	repository "github.com/pedro-muniz/myPrice/src/product-service/core/port/repository"
	port "github.com/pedro-muniz/myPrice/src/product-service/core/port/usecase/price"
)

type PriceManagement struct {
	PriceRepository repository.IPriceRepository
}

var priceManagementOnce sync.Once
var priceManagementInstance *PriceManagement

func GetPriceManagementInstance(priceRepository repository.IPriceRepository) *PriceManagement {
	priceManagementOnce.Do(func() {
		priceManagementInstance = &PriceManagement{
			PriceRepository: priceRepository,
		}
	})

	return priceManagementInstance
}

func (this *PriceManagement) Create(input *port.CreateInput) (*port.CreateOutput, error) {
	if input == nil {
		return nil, priceErrors.InvalidPriceData("nil input")
	}

	price := domain.NewPrice(
		input.Gross,
		input.Net,
		input.Selling,
		input.Recommended,
	)
	price.CreatedAt = time.Now()
	price.UpdatedAt = time.Now()

	if err := price.Validate(); err != nil {
		return nil, priceErrors.InvalidPriceData(err.Error())
	}

	priceChan, errChan := this.PriceRepository.Save(price)

	var savedPrice *domain.Price
	var err error

	select {
	case savedPrice = <-priceChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, priceErrors.ErrorSavingPriceDatabaseRecord(err.Error())
	}

	if savedPrice == nil {
		return nil, priceErrors.ErrorSavingPriceDatabaseRecord("saved price is nil")
	}

	return &port.CreateOutput{
		Id: savedPrice.Id,
	}, nil
}

func (this *PriceManagement) Get(input *port.GetInput) (*port.GetOutput, error) {
	if input == nil || input.Id == "" {
		return nil, priceErrors.InvalidPriceData("invalid id")
	}

	priceChan, errChan := this.PriceRepository.FindById(input.Id)

	var price *domain.Price
	var err error

	select {
	case price = <-priceChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, priceErrors.ErrorGettingPriceDatabaseRecord(err.Error())
	}

	if price == nil {
		return nil, priceErrors.PriceNotFound(input.Id)
	}

	return &port.GetOutput{
		Price: price,
	}, nil
}

func (this *PriceManagement) List(input *port.ListInput) (*port.ListOutput, error) {
	pricesChan, errChan := this.PriceRepository.FindAll()

	var prices []*domain.Price
	var err error

	select {
	case prices = <-pricesChan:
	case err = <-errChan:
	}

	if err != nil {
		return nil, priceErrors.ErrorGettingPriceDatabaseRecord(err.Error())
	}

	return &port.ListOutput{
		Prices: prices,
	}, nil
}

func (this *PriceManagement) Update(input *port.UpdateInput) (*port.UpdateOutput, error) {
	if input == nil || input.Id == "" {
		return nil, priceErrors.InvalidPriceData("invalid id")
	}

	priceChan, findErrChan := this.PriceRepository.FindById(input.Id)
	var existingPrice *domain.Price
	var findErr error

	select {
	case existingPrice = <-priceChan:
	case findErr = <-findErrChan:
	}

	if findErr != nil {
		return nil, priceErrors.ErrorGettingPriceDatabaseRecord(findErr.Error())
	}
	if existingPrice == nil {
		return nil, priceErrors.PriceNotFound(input.Id)
	}

	existingPrice.Gross = input.Gross
	existingPrice.Net = input.Net
	existingPrice.Selling = input.Selling
	existingPrice.Recommended = input.Recommended
	existingPrice.UpdatedAt = time.Now()

	if err := existingPrice.Validate(); err != nil {
		return nil, priceErrors.InvalidPriceData(err.Error())
	}

	errChan := this.PriceRepository.Update(existingPrice)

	var err error = <-errChan
	if err != nil {
		return nil, priceErrors.ErrorSavingPriceDatabaseRecord(err.Error())
	}

	return &port.UpdateOutput{
		Success: true,
	}, nil
}

func (this *PriceManagement) Delete(input *port.DeleteInput) (*port.DeleteOutput, error) {
	if input == nil || input.Id == "" {
		return nil, priceErrors.InvalidPriceData("invalid id")
	}

	errChan := this.PriceRepository.Delete(input.Id)

	var err error = <-errChan
	if err != nil {
		return nil, priceErrors.ErrorSavingPriceDatabaseRecord(err.Error())
	}

	return &port.DeleteOutput{
		Success: true,
	}, nil
}

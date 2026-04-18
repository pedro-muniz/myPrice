package price

import (
	"fmt"
)

type InvalidPriceData string

func (e InvalidPriceData) Error() string {
	return fmt.Sprintf("Invalid price data: %s", string(e))
}

type PriceNotFound string

func (e PriceNotFound) Error() string {
	return fmt.Sprintf("Price not found: %s", string(e))
}

type ErrorGettingPriceDatabaseRecord string

func (e ErrorGettingPriceDatabaseRecord) Error() string {
	return fmt.Sprintf("Error getting price record: %s", string(e))
}

type ErrorSavingPriceDatabaseRecord string

func (e ErrorSavingPriceDatabaseRecord) Error() string {
	return fmt.Sprintf("Error saving price record: %s", string(e))
}

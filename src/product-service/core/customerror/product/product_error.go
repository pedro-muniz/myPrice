package product

import (
	"fmt"
)

type InvalidProductData string

func (e InvalidProductData) Error() string {
	return fmt.Sprintf("Invalid product data: %s", string(e))
}

type ProductNotFound string

func (e ProductNotFound) Error() string {
	return fmt.Sprintf("Product not found: %s", string(e))
}

type ErrorGettingProductDatabaseRecord string

func (e ErrorGettingProductDatabaseRecord) Error() string {
	return fmt.Sprintf("Error getting product record: %s", string(e))
}

type ErrorSavingProductDatabaseRecord string

func (e ErrorSavingProductDatabaseRecord) Error() string {
	return fmt.Sprintf("Error saving product record: %s", string(e))
}

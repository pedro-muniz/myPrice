package domain

import (
	"errors"
	"time"
)

type Origin int

const (
	National Origin = iota
	Imported
)

type Product struct {
	Id            string
	CompanyId     string
	BranchId      string
	BarCode       string
	Name          string
	Description   string
	Discount      float64
	Price         Price
	ProfitPercent float64
	Origin        Origin
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

func NewProduct(companyId string, branchId string, barCode string,
	name string, description string,
	discount float64, price Price, profitPercent float64,
	origin Origin) *Product {

	return &Product{
		CompanyId:     companyId,
		BranchId:      branchId,
		BarCode:       barCode,
		Name:          name,
		Description:   description,
		Discount:      discount,
		Price:         price,
		ProfitPercent: profitPercent,
		Origin:        origin,
	}
}

func (this *Product) Validate() error {
	if len(this.CompanyId) <= 0 {
		return errors.New("invalid product company id")
	}

	if len(this.BranchId) <= 0 {
		return errors.New("invalid product branch id")
	}

	if len(this.BarCode) <= 0 {
		return errors.New("invalid product barcode")
	}

	if len(this.Name) <= 0 {
		return errors.New("invalid product name")
	}

	if len(this.Description) <= 0 {
		return errors.New("invalid product description")
	}

	if this.Origin < National || this.Origin > Imported {
		return errors.New("invalid product origin")
	}

	if this.Price.Validate() != nil {
		return errors.New("invalid product price")
	}

	if this.CreatedAt.IsZero() {
		return errors.New("invalid product created at")
	}

	if this.UpdatedAt.IsZero() {
		return errors.New("invalid product updated at")
	}

	return nil
}
func (this *Product) Delete() {
	now := time.Now()
	this.DeletedAt = &now
	this.Price.Delete()
}

package domain

import (
	"errors"
	"time"
)

type Price struct {
	Id          string
	CompanyId   string
	BranchId    string
	Gross       float64
	Net         float64
	Selling     float64
	Recommended float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func NewPrice(companyId string, branchId string, gross float64,
	net float64, selling float64, recommended float64) *Price {
	return &Price{
		CompanyId:   companyId,
		BranchId:    branchId,
		Gross:       gross,
		Net:         net,
		Selling:     selling,
		Recommended: recommended,
	}
}

func (this *Price) Validate() error {
	if len(this.CompanyId) <= 0 {
		return errors.New("invalid price company id")
	}

	if len(this.BranchId) <= 0 {
		return errors.New("invalid price branch id")
	}

	if this.Gross <= 0 {
		return errors.New("invalid price gross")
	}

	if this.Net <= 0 {
		return errors.New("invalid price net")
	}

	if this.Selling <= 0 {
		return errors.New("invalid price selling")
	}

	if this.Recommended <= 0 {
		return errors.New("invalid price recommended")
	}

	if this.CreatedAt.IsZero() {
		return errors.New("invalid price created at")
	}

	if this.UpdatedAt.IsZero() {
		return errors.New("invalid price updated at")
	}

	return nil
}
func (this *Price) Delete() {
	now := time.Now()
	this.DeletedAt = &now
}

package domain

import (
	"errors"
	"time"
)

type Price struct {
	Id          string
	Gross       float64
	Net         float64
	Selling     float64
	Recommended float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewPrice(gross float64, net float64, selling float64, recommended float64) *Price {
	return &Price{
		Gross:       gross,
		Net:         net,
		Selling:     selling,
		Recommended: recommended,
	}
}

func (this *Price) Validate() error {
	if len(this.Id) <= 0 {
		return errors.New("invalid price id")
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

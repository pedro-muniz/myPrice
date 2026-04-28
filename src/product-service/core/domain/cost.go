package domain

import (
	"errors"
	"time"
)

type Cost struct {
	Id        string
	CompanyId string
	BranchId  string
	Name      string
	Value     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewCost(id string, companyId string, branchId string, name string, value float64) *Cost {
	now := time.Now()
	return &Cost{
		Id:        id,
		CompanyId: companyId,
		BranchId:  branchId,
		Name:      name,
		Value:     value,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (this *Cost) Validate() error {
	if len(this.CompanyId) <= 0 {
		return errors.New("invalid cost company id")
	}

	if len(this.BranchId) <= 0 {
		return errors.New("invalid cost branch id")
	}

	if len(this.Name) <= 0 {
		return errors.New("invalid cost name")
	}

	if this.Value < 0 {
		return errors.New("invalid cost value")
	}

	return nil
}

func (this *Cost) Delete() {
	now := time.Now()
	this.DeletedAt = &now
}

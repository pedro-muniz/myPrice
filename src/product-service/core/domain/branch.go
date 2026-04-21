package domain

import (
	"errors"
	"time"
)

type Branch struct {
	Id        string
	CompanyId string
	Name      string
	DeletedAt *time.Time
}

func NewBranch(id string, companyId string, name string) *Branch {
	return &Branch{
		Id:        id,
		CompanyId: companyId,
		Name:      name,
	}
}

func (this *Branch) Validate() error {
	if len(this.CompanyId) <= 0 {
		return errors.New("invalid company id")
	}

	if len(this.Name) <= 0 {
		return errors.New("invalid branch name")
	}

	return nil
}
func (this *Branch) Delete(deletedAt time.Time) {
	this.DeletedAt = &deletedAt
}

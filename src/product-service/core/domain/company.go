package domain

import (
	"errors"
	"time"
)

type Company struct {
	Id        string
	Name      string
	DeletedAt *time.Time
}

func NewCompany(id string, name string) *Company {
	return &Company{
		Id:   id,
		Name: name,
	}
}

func (this *Company) Validate() error {
	if len(this.Name) <= 0 {
		return errors.New("invalid company name")
	}

	return nil
}
func (this *Company) Delete(deletedAt time.Time) {
	this.DeletedAt = &deletedAt
}

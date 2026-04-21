package company

import (
	"fmt"
)

type InvalidCompanyData string

func (e InvalidCompanyData) Error() string {
	return fmt.Sprintf("Invalid company data: %s", string(e))
}

type CompanyNotFound string

func (e CompanyNotFound) Error() string {
	return fmt.Sprintf("Company not found: %s", string(e))
}

type ErrorGettingCompanyDatabaseRecord string

func (e ErrorGettingCompanyDatabaseRecord) Error() string {
	return fmt.Sprintf("Error getting company record: %s", string(e))
}

type ErrorSavingCompanyDatabaseRecord string

func (e ErrorSavingCompanyDatabaseRecord) Error() string {
	return fmt.Sprintf("Error saving company record: %s", string(e))
}

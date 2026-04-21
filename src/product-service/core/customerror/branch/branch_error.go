package branch

import (
	"fmt"
)

type InvalidBranchData string

func (e InvalidBranchData) Error() string {
	return fmt.Sprintf("Invalid branch data: %s", string(e))
}

type BranchNotFound string

func (e BranchNotFound) Error() string {
	return fmt.Sprintf("Branch not found: %s", string(e))
}

type ErrorGettingBranchDatabaseRecord string

func (e ErrorGettingBranchDatabaseRecord) Error() string {
	return fmt.Sprintf("Error getting branch record: %s", string(e))
}

type ErrorSavingBranchDatabaseRecord string

func (e ErrorSavingBranchDatabaseRecord) Error() string {
	return fmt.Sprintf("Error saving branch record: %s", string(e))
}

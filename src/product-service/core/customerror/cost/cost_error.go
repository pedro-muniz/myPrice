package cost

import (
	"fmt"
)

type InvalidCostData string

func (e InvalidCostData) Error() string {
	return fmt.Sprintf("Invalid cost data: %s", string(e))
}

type CostNotFound string

func (e CostNotFound) Error() string {
	return fmt.Sprintf("Cost not found: %s", string(e))
}

type ErrorGettingCostDatabaseRecord string

func (e ErrorGettingCostDatabaseRecord) Error() string {
	return fmt.Sprintf("Error getting cost record: %s", string(e))
}

type ErrorSavingCostDatabaseRecord string

func (e ErrorSavingCostDatabaseRecord) Error() string {
	return fmt.Sprintf("Error saving cost record: %s", string(e))
}

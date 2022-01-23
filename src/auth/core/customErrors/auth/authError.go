package errors

import (
	"fmt"
)

// Invalid auth clientId
type InvalidClientId string

func (f InvalidClientId) Error() string {
	return "Request had invalid authentication credentials."
}

// Invalid auth object reference.
type InvalidAuthReference string

func (f InvalidAuthReference) Error() string {
	return "Request had invalid authentication reference."
}

// Error getting auth database record
type ErrorGettingAuthDatabaseRecord string

func (f ErrorGettingAuthDatabaseRecord) Error() string {
	return fmt.Sprintf("Error getting client credentials. %s", f.Error())
}

// Client not found on database
type ClientNotFound string

func (f ClientNotFound) Error() string {
	return fmt.Sprintf("Invalid client credentials.")
}

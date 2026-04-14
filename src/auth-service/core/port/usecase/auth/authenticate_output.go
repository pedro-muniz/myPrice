package auth

import "time"

// AuthenticateOutput is the output DTO returned by the Authenticate use case
type AuthenticateOutput struct {
	ClientId   string
	Token      string
	ExpiringIn time.Duration
}

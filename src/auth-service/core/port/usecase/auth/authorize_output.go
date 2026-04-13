package auth

import "time"

// AuthorizeOutput is the output DTO returned by the Authorize use case
type AuthorizeOutput struct {
	ClientId   string
	Token      string
	ExpiringIn time.Duration
}

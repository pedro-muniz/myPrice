package auth

// AuthenticateInput is the input DTO for the Authenticate use case
type AuthenticateInput struct {
	ClientId     string
	ClientSecret string
}

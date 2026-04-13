package auth

// IAuthenticate defines the behavior for the authentication use case
type IAuthenticate interface {
	Execute(input *AuthenticateInput) (*AuthenticateOutput, error)
}

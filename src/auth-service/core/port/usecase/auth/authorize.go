package auth

// IAuthorize defines the behavior for the authorization use case (JWT token validation)
type IAuthorize interface {
	Execute(token string) (*AuthorizeOutput, error)
}

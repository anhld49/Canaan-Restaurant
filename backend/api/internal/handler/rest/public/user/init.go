package user

import "backend/api/internal/controller/user"

// AuthenticationRequestPayload: Authentication infor from request payload
type AuthenticationRequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// EmailRequestPayload: read email from request payload
type EmailRequestPayload struct {
	Email string `json:"email"`
}

// IdRequestPayload: read id from request payload
type IdRequestPayload struct {
	Id int `json:"id"`
}

// UserHandler: User handler
type UserHandler struct {
	controller user.UserController
}

// NewUserHandler: Create new user handler
func NewUserHandler(handler user.UserController) *UserHandler {
	return &UserHandler{
		controller: handler,
	}
}

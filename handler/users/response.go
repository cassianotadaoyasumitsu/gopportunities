package users

import "gopportunities/schemas"

type SignupUserResponse struct {
	Message string               `json:"message"`
	Data    schemas.UserResponse `json:"data"`
}

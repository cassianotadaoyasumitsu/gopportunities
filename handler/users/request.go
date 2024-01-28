package users

import "gopportunities/handler"

type SignUpUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInUserRequest struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *SignUpUserRequest) Validate() error {
	if r.Email == "" && r.Password == "" {
		return handler.ErrParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return handler.ErrParamIsRequired("password", "string")
	}
	return nil
}

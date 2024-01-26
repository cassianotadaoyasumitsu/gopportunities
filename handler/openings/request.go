package openings

import (
	"fmt"
	"gopportunities/handler"
)

//func errParamIsRequired(name, typ string) error {
//	return fmt.Errorf("param: %s (type%s) is required", name, typ)
//}

// CreateOpeningRequest is the request body for CreateOpeningHandler
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Salary <= 0 && r.Remote == nil {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return handler.ErrParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return handler.ErrParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return handler.ErrParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return handler.ErrParamIsRequired("link", "string")
	}
	if r.Salary <= 0 {
		return handler.ErrParamIsRequired("salary", "int64")
	}
	if r.Remote == nil {
		return handler.ErrParamIsRequired("remote", "bool")
	}
	return nil
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" && r.Company != "" && r.Location != "" && r.Link != "" && r.Salary > 0 && r.Remote != nil {
		return nil
	}

	return fmt.Errorf("at least one field is required")
}

package transport

import (
	"bootcampProject/users/domain"
	"gopkg.in/validator.v2"
)

// CreateUserRequest holds the request parameters for the CreateUser method.
type CreateUserRequest struct {
	User domain.Users
}

func (req *CreateUserRequest) Validate() error {
	return validator.Validate(req.User)
}

// CreateUserResponse holds the response values for the CreateUser method.
type CreateUserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Err   error  `json:"error,omitempty"`
}

// GetUsersRequest holds the request parameters for the GetUsers method.
type GetUsersRequest struct {
	limit  int `validate:"min= 1"`
	offset int `validate:"min= 0"`
}

func (req *GetUsersRequest) Validate() error {
	return validator.Validate(req)
}

// GetUsersResponse holds the response values for the GetUsers method.
type GetUsersResponse struct {
	Users []domain.Users `json:"users"`
	Err   error          `json:"error,omitempty"`
}

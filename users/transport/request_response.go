package transport

import "bootcampProject/users/domain"

// CreateUserRequest holds the request parameters for the CreateUser method.
type CreateUserRequest struct {
	User domain.Users
}

// CreateUserResponse holds the response values for the CreateUser method.
type CreateUserResponse struct {
	ID  int   `json:"id"`
	Err error `json:"error,omitempty"`
}

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

// GetUsersRequest holds the request parameters for the GetUsers method.
type GetUsersRequest struct {
	limit  int
	offset int
}

// GetUsersResponse holds the response values for the GetUsers method.
type GetUsersResponse struct {
	Users []domain.Users `json:"users"`
	Err   error          `json:"error,omitempty"`
}

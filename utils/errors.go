package utils

import "errors"

const DescriptionBadRequest = "bad request. It doesn't match interface"
const DescriptionInvalidCredentials = "invalid credentials"

type ErrBadRequest struct {
	error
}

type ErrInvalidCredentials struct {
	error
}

func NewErrBadRequest() *ErrBadRequest {
	return &ErrBadRequest{error: errors.New(DescriptionBadRequest)}
}

func NewErrInvalidCredentials() *ErrInvalidCredentials {
	return &ErrInvalidCredentials{error: errors.New(DescriptionInvalidCredentials)}
}

var (
	ErrRecordNotFound = errors.New("record not found")
)

package utils

import "errors"

var (
	ErrBadRequest         = errors.New("bad request. It doesn't match interface")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrRecordNotFound     = errors.New("record not found")
)

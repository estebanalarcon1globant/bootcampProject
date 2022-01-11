package http

import (
	"bootcampProject/users/transport"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

// NewUserHandler wires Go kit endpoints to the HTTP transport.
func NewUserHandler(
	svcEndpoints transport.UserEndpoints, logger log.Logger) http.Handler {
	// set-up router and initialize http endpoints
	r := mux.NewRouter()
	// HTTP Post - /orders
	r.Methods("POST").Path("/users").Handler(kithttp.NewServer(
		svcEndpoints.CreateUser,
		decodeCreateUserRequest,
		encodeResponse,
	))
	return r
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.CreateUserRequest
	if e := json.NewDecoder(r.Body).Decode(&req.User); e != nil {
		return nil, e
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

/*
type errorer interface {
	error() error
}*/

/*
func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}*/

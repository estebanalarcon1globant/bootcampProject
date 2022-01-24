package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

// NewUserHTTPServer wires Go kit endpoints to the HTTP transport.
func NewUserHTTPServer(svcEndpoints UserEndpointsHTTP, logger log.Logger) http.Handler {
	// set-up router and initialize http endpoints
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		//kithttp.ServerErrorEncoder(encodeError),
	}
	// HTTP Post - /users
	r.Methods("POST").Path("/users").Handler(kithttp.NewServer(
		svcEndpoints.CreateUser,
		decodeCreateUserHTTPRequest,
		encodeCreateUserHTTPResponse,
		opts...,
	))

	// HTTP Get - /users
	r.Methods("GET").Path("/users").Handler(kithttp.NewServer(
		svcEndpoints.GetUsers,
		decodeGetUsersHTTPRequest,
		encodeGetUsersHTTPResponse,
		opts...,
	))
	return r
}

func decodeCreateUserHTTPRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req CreateUserRequest
	if e := json.NewDecoder(r.Body).Decode(&req.User); e != nil {
		return nil, e
	}
	return req, nil
}

func encodeCreateUserHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeGetUsersHTTPRequest(_ context.Context, r *http.Request) (request interface{}, err error) {

	vars := r.URL.Query()
	var limit, offset int
	if auxLimit := vars.Get("limit"); auxLimit == "" {
		limit = 100
	} else {
		limit, err = strconv.Atoi(auxLimit)
	}

	if auxOffset := vars.Get("offset"); auxOffset == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(auxOffset)
	}

	return GetUsersRequest{
		limit:  limit,
		offset: offset,
	}, err
}

func encodeGetUsersHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

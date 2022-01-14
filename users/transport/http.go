package transport

import (
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

// NewUserHTTPServer wires Go kit endpoints to the HTTP transport.
func NewUserHTTPServer(
	svcEndpoints UserEndpointsHTTP, logger log.Logger) http.Handler {
	// set-up router and initialize http endpoints
	r := mux.NewRouter()
	// HTTP Post - /orders
	r.Methods("POST").Path("/users").Handler(kithttp.NewServer(
		svcEndpoints.CreateUser,
		decodeCreateUserHTTPRequest,
		encodeCreateUserHTTPResponse,
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

func encodeCreateUserHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

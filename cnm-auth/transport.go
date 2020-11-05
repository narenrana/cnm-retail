package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the booking service.
func MakeHandler(bs Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	authLoginHandler := kithttp.NewServer(
		makeAuthLoginEndpoint(bs),
		decodeAuthLoginRequest,
		encodeResponse,
		opts...,
	)

	authLogoutHandler := kithttp.NewServer(
		makeAuthLogoutEndpoint(bs),
		decodeAuthLogoutRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/auth/v1/login", authLoginHandler).Methods("POST")
	r.Handle("/auth/v1/logout", authLogoutHandler).Methods("GET")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeAuthLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name       string `json:"name"`
		Password   string `json:"password"`
		RememberMe string `json:"rememberMe"`
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&body); err != nil {
		return nil, err
	}

	return authLoginRequest{
		name:       body.Name,
		password:   body.Password,
		rememberMe: body.RememberMe,
	}, nil
}

func decodeAuthLogoutRequest(_ context.Context, r *http.Request) (interface{}, error) {

	value := r.FormValue("token")
	return authLogoutRequest{
		token:       value,
	}, nil
}


func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

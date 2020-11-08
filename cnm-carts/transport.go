package carts

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"shopping-cart/cnm-carts/models"
	"shopping-cart/cnm-carts/services"
	"strconv"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the booking service.
func MakeHandler(bs services.Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	addToCartHandler := kithttp.NewServer(
		makeAddToCartEndpoint(bs),
		decodeAddToCartRequest,
		encodeResponse,
		opts...,
	)

	getCartHandler := kithttp.NewServer(
		makeGetCartEndpoint(bs),
		decodeGetCartRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/carts/v1/add", addToCartHandler).Methods("POST")
	r.Handle("/carts/v1/list", getCartHandler).Methods("GET")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeAddToCartRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request models.AddToCartRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetCartRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var getCartRequest models.GetCartRequest;
	value := r.FormValue("userId")
	userId, err :=strconv.ParseInt(value,10, 64)
	getCartRequest.UseId=int(userId)
	return getCartRequest, err
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
	case services.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeCreateEndpoint(s OrderService) endpoint.Endpoint {
	fmt.Println("createEndpoint called")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		id, err := s.Create(ctx, req.Order)
		return CreateResponse{ID: id, Err: err}, nil
	}
}

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("decode requests called")
	var request CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&request.Order); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("encodeResponse Called")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type CreateRequest struct {
	Order Order
}

type CreateResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

//

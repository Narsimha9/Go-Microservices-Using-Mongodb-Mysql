package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

// Endpoint for the Account service.

func makeCreateCustomerEndpoint(s AccountService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCustomerRequest)
		id, err := s.CreateCustomer(ctx, req.customer)
		return CreateCustomerResponse{Id: id, Err: err}, nil
	}

}

func makeGetCustomerByIdEndpoint(s AccountService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCustomerByIdRequest)
		fmt.Println("Request", req)
		id, er := strconv.Atoi(req.Id)
		if er != nil {
			return GetCustomerByIdResponse{Data: "", Err: er}, nil
		}
		data, err := s.GetCustomerById(ctx, id)
		fmt.Println("ID decoded output:", data)
		return GetCustomerByIdResponse{Data: data, Err: err}, nil
	}

}
func makeGetAllCustomersEndpoint(s AccountService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		data, err := s.GetAllCustomers(ctx)
		return GetAllCustomersResponse{Data: data, Err: err}, nil
	}

}
func makeDeleteCustomerEndpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteCustomerRequest)
		fmt.Println("Request of DeleteCustomer", req)
		fmt.Println("Rquest ud:", req.Id)
		id, er := strconv.Atoi(req.Id)
		if er != nil {
			return DeleteCustomerResponse{Msg: "", Err: er}, nil
		}
		msg, err := s.DeleteCustomer(ctx, id)
		return DeleteCustomerResponse{Msg: msg, Err: err}, nil
	}

}

func makeUpdateCustomerEndpoint(s AccountService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateCustomerRequest)
		fmt.Println("Request", req.customer)
		rc := req.customer
		fmt.Println("Request Id", rc.Customerid)
		fmt.Println("REQ.customer:", req.customer)
		err := s.UpdateCustomer(ctx, rc.Customerid, req.customer)
		return UpdateCustomerResponse{Err: err}, nil
	}
}

func decodeCreateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req.customer); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetCustomerByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetCustomerByIdRequest
	vars := mux.Vars(r)
	req = GetCustomerByIdRequest{
		Id: vars["id"],
	}

	return req, nil
}

func decodeGetAllCustomersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetAllCustomersRequest

	return req, nil
}

func decodeDeleteCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req DeleteCustomerRequest
	vars := mux.Vars(r)
	req = DeleteCustomerRequest{
		Id: vars["id"],
	}

	return req, nil
}

func decodeUpdateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req UpdateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req.customer); err != nil {
		return nil, err
	}
	return req, nil

}

//  encodes the output
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type (
	CreateCustomerRequest struct {
		customer Customer
	}
	CreateCustomerResponse struct {
		Id  string `json:"id"`
		Err error
	}
	GetCustomerByIdRequest struct {
		Id string `json:"id"`
	}
	GetCustomerByIdResponse struct {
		Data interface{} `json:"customer"`
		Err  error       `json:"error,omitempty"`
	}
	GetAllCustomersRequest struct {
	}
	GetAllCustomersResponse struct {
		Data interface{} `json:"customer"`
		Err  error       `json:"error,omitempty"`
	}
	DeleteCustomerRequest struct {
		Id string `json:"id"`
	}
	DeleteCustomerResponse struct {
		Msg string `json:"msg,omitempty"`
		Err error  `json:"error,omitempty"`
	}
	UpdateCustomerRequest struct {
		ID       string `json:"id"`
		customer Customer
	}
	UpdateCustomerResponse struct {
		Err error `json:"error,omitempty"`
	}
)

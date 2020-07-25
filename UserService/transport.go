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

// Endpoint for the User service.

func makeCreateUserEndpoint(s UserService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		id, err := s.CreateUser(ctx, req.user)
		return CreateUserResponse{Id: id, Err: err}, nil
	}

}

func makeGetUserByIdEndpoint(s UserService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserByIdRequest)
		fmt.Println("Request", req)
		id, er := strconv.Atoi(req.Id)
		if er != nil {
			return GetUserByIdResponse{Data: "", Err: er}, nil
		}
		data, err := s.GetUserById(ctx, id)
		fmt.Println("ID decoded output:", data)
		return GetUserByIdResponse{Data: data, Err: err}, nil
	}

}
func makeGetAllUsersEndpoint(s UserService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		data, err := s.GetAllUsers(ctx)
		return GetAllUsersResponse{Data: data, Err: err}, nil
	}

}
func makeDeleteUserEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		fmt.Println("Request of DeleteUser", req)
		fmt.Println("Rquest ud:", req.Id)
		id, er := strconv.Atoi(req.Id)
		if er != nil {
			return DeleteUserResponse{Msg: "", Err: er}, nil
		}
		msg, err := s.DeleteUser(ctx, id)
		return DeleteUserResponse{Msg: msg, Err: err}, nil
	}

}

func makeUpdateUserEndpoint(s UserService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		fmt.Println("Request", req.user)
		rc := req.user
		fmt.Println("Request Id", rc.Userid)
		fmt.Println("REQ.user:", req.user)
		err := s.UpdateUser(ctx, rc.Userid, req.user)
		return UpdateUserResponse{Err: err}, nil
	}
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req.user); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetUserByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetUserByIdRequest
	vars := mux.Vars(r)
	req = GetUserByIdRequest{
		Id: vars["id"],
	}

	return req, nil
}

func decodeGetAllUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetAllUsersRequest

	return req, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req DeleteUserRequest
	vars := mux.Vars(r)
	req = DeleteUserRequest{
		Id: vars["id"],
	}

	return req, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req.user); err != nil {
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
	CreateUserRequest struct {
		user User
	}
	CreateUserResponse struct {
		Id  string `json:"id"`
		Err error
	}
	GetUserByIdRequest struct {
		Id string `json:"id"`
	}
	GetUserByIdResponse struct {
		Data interface{} `json:"user"`
		Err  error       `json:"error,omitempty"`
	}
	GetAllUsersRequest struct {
	}
	GetAllUsersResponse struct {
		Data interface{} `json:"user"`
		Err  error       `json:"error,omitempty"`
	}
	DeleteUserRequest struct {
		Id string `json:"id"`
	}
	DeleteUserResponse struct {
		Msg string `json:"msg,omitempty"`
		Err error  `json:"error,omitempty"`
	}
	UpdateUserRequest struct {
		ID   string `json:"id"`
		user User
	}
	UpdateUserResponse struct {
		Err error `json:"error,omitempty"`
	}
)

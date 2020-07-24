package main

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

// service implements the ACcount Service
type accountservice struct {
	repository Repository
	logger     log.Logger
}

// Service describes the Account service.
type AccountService interface {
	CreateCustomer(ctx context.Context, customer Customer) (string, error)
	GetCustomerById(ctx context.Context, id int) (interface{}, error)
	GetAllCustomers(ctx context.Context) (interface{}, error)
	DeleteCustomer(ctx context.Context, id int) (string, error)
	UpdateCustomer(ctx context.Context, id int, customer Customer) error
}

// NewService creates and returns a new Account service instance
func NewService(rep Repository, logger log.Logger) AccountService {
	return &accountservice{
		repository: rep,
		logger:     logger,
	}
}

// Create makes an customer
func (s accountservice) CreateCustomer(ctx context.Context, customer Customer) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	uuid, _ := uuid.NewV4()
	var id = uuid.String()
	customer.Id = id
	customerDetails := Customer{
		Id:         customer.Id,
		Customerid: customer.Customerid,
		Email:      customer.Email,
		Password:   customer.Password,
		Phone:      customer.Phone,
	}
	if err := s.repository.CreateCustomer(ctx, customerDetails); err != nil {
		level.Error(logger).Log("err", err)
	}
	return id, nil
}

func (s accountservice) GetCustomerById(ctx context.Context, id int) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetCustomerById")
	var data interface{}
	data, err := s.repository.GetCustomerById(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)

		return "", err
	}
	return data, nil
}

func (s accountservice) GetAllCustomers(ctx context.Context) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetAllCustomers")
	var email interface{}
	email, err := s.repository.GetAllCustomers(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	return email, nil
}

func (s accountservice) DeleteCustomer(ctx context.Context, id int) (string, error) {
	logger := log.With(s.logger, "method", "DeleteCustomer")
	msg, err := s.repository.DeleteCustomer(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)

		return "", err
	}
	return msg, nil
}

func (s accountservice) UpdateCustomer(ctx context.Context, id int, customer Customer) error {
	logger := log.With(s.logger, "method", "ChangeDetails")
	// var email string
	err := s.repository.UpdateCustomer(ctx, id, customer)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	} else {
		// msg := "Data Updated Successfully"
		return nil
	}

}

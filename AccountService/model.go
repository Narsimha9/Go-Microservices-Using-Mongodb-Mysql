package main

import "context"

type Customer struct {
	Id         string `json:"id"`
	Customerid int    `json:"customerid"`
	Email      string ` json:"email"`
	Password   string ` json:"password"`
	Phone      string ` json:"phone"`
}

type Repository interface {
	CreateCustomer(ctx context.Context, customer Customer) error
	GetCustomerById(ctx context.Context, id int) (interface{}, error)
	GetAllCustomers(ctx context.Context) (interface{}, error)
	DeleteCustomer(ctx context.Context, id int) (string, error)
	UpdateCustomer(ctx context.Context, id int, customer Customer) error
}

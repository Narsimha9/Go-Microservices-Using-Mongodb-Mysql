package main

import "context"

type User struct {
	Id         string `json:"id"`
	Userid int    `json:"userid"`
	Email      string ` json:"email"`
	Password   string ` json:"password"`
	Phone      string ` json:"phone"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUserById(ctx context.Context, id int) (interface{}, error)
	GetAllUsers(ctx context.Context) (interface{}, error)
	DeleteUser(ctx context.Context, id int) (string, error)
	UpdateUser(ctx context.Context, id int, user User) error
}

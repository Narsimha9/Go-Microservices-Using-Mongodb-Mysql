package main

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

// service implements the ACcount Service
type userserviceStruct struct {
	repository Repository
	logger     log.Logger
}

// Service describes the User service.
type UserService interface {
	CreateUser(ctx context.Context, user User) (string, error)
	GetUserById(ctx context.Context, id int) (interface{}, error)
	GetAllUsers(ctx context.Context) (interface{}, error)
	DeleteUser(ctx context.Context, id int) (string, error)
	UpdateUser(ctx context.Context, id int, user User) error
}

// NewService creates and returns a new User service instance
func NewService(rep Repository, logger log.Logger) UserService {
	return &userserviceStruct{
		repository: rep,
		logger:     logger,
	}
}

// Create makes an user
func (s userserviceStruct) CreateUser(ctx context.Context, user User) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	uuid, _ := uuid.NewV4()
	var id = uuid.String()
	user.Id = id
	userDetails := User{
		Id:       user.Id,
		Userid:   user.Userid,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
	}
	if err := s.repository.CreateUser(ctx, userDetails); err != nil {
		level.Error(logger).Log("err", err)
	}
	return id, nil
}

func (s userserviceStruct) GetUserById(ctx context.Context, id int) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetUserById")
	var data interface{}
	data, err := s.repository.GetUserById(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)

		return "", err
	}
	return data, nil
}

func (s userserviceStruct) GetAllUsers(ctx context.Context) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetAllUsers")
	var email interface{}
	email, err := s.repository.GetAllUsers(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	return email, nil
}

func (s userserviceStruct) DeleteUser(ctx context.Context, id int) (string, error) {
	logger := log.With(s.logger, "method", "DeleteUser")
	msg, err := s.repository.DeleteUser(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)

		return "", err
	}
	return msg, nil
}

func (s userserviceStruct) UpdateUser(ctx context.Context, id int, user User) error {
	logger := log.With(s.logger, "method", "ChangeDetails")
	// var email string
	err := s.repository.UpdateUser(ctx, id, user)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	} else {
		// msg := "Data Updated Successfully"
		return nil
	}

}

package main

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

func (mw loggingMiddleware) CreateUser(ctx context.Context, user User) (Email string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "creatuser",
			"Email", user.Email,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	Email, err = mw.next.CreateUser(ctx, user)
	return
}

func (mw loggingMiddleware) GetUserById(ctx context.Context, id int) (data interface{}, Err error) {
	defer func(begin time.Time) {

		_ = mw.logger.Log(
			"method", "GetbyId",
			"id", id,
			"err", Err,
			"took", time.Since(begin),
		)
	}(time.Now())

	data, Err = mw.next.GetUserById(ctx, id)
	return
}

func (mw loggingMiddleware) GetAllUsers(ctx context.Context) (Email interface{}, Err error) {
	defer func(begin time.Time) {

		_ = mw.logger.Log(
			"method", "GetAllUsers",
			"err", Err,
			"took", time.Since(begin),
		)
	}(time.Now())

	Email, Err = mw.next.GetAllUsers(ctx)
	return
}
func (mw loggingMiddleware) DeleteUser(ctx context.Context, id int) (Msg string, Err error) {
	defer func(begin time.Time) {

		_ = mw.logger.Log(
			"method", "DeleteUser",
			"id", id,
			"err", Err,
			"took", time.Since(begin),
		)
	}(time.Now())

	Msg, Err = mw.next.DeleteUser(ctx, id)

	return
}

func (mw loggingMiddleware) UpdateUser(ctx context.Context, id int, user User) (Err error) {
	defer func(begin time.Time) {

		_ = mw.logger.Log(
			"method", "update_user",
			"Msg", "Data updated",
			"err", Err,
			"took", time.Since(begin),
		)
	}(time.Now())

	Err = mw.next.UpdateUser(ctx, id, user)
	return
}

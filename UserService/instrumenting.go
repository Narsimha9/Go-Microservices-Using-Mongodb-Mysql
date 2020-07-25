package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           UserService
}

func (mw instrumentingMiddleware) CreateUser(ctx context.Context, user User) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "createuser", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.CreateUser(ctx, user)
	return
}

func (mw instrumentingMiddleware) GetUserById(ctx context.Context, id int) (data interface{}, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetUserById", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	data, err = mw.next.GetUserById(ctx, id)
	return
}

func (mw instrumentingMiddleware) GetAllUsers(ctx context.Context) (Email interface{}, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAllUsers", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	Email, err = mw.next.GetAllUsers(ctx)
	return
}
func (mw instrumentingMiddleware) DeleteUser(ctx context.Context, id int) (msg string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "DeleteUser", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	msg, err = mw.next.DeleteUser(ctx, id)
	return
}
func (mw instrumentingMiddleware) UpdateUser(ctx context.Context, id int, user User) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateUser", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	err = mw.next.UpdateUser(ctx, id, user)
	return
}

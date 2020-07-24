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
	next           AccountService
}

func (mw instrumentingMiddleware) CreateCustomer(ctx context.Context, customer Customer) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "createcustomer", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.CreateCustomer(ctx, customer)
	return
}

func (mw instrumentingMiddleware) GetCustomerById(ctx context.Context, id int) (data interface{}, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetCustomerById", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	data, err = mw.next.GetCustomerById(ctx, id)
	return
}

func (mw instrumentingMiddleware) GetAllCustomers(ctx context.Context) (Email interface{}, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAllCustomers", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	Email, err = mw.next.GetAllCustomers(ctx)
	return
}
func (mw instrumentingMiddleware) DeleteCustomer(ctx context.Context, id int) (msg string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "DeleteCustomer", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	msg, err = mw.next.DeleteCustomer(ctx, id)
	return
}
func (mw instrumentingMiddleware) UpdateCustomer(ctx context.Context, id int, customer Customer) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateCustomer", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	err = mw.next.UpdateCustomer(ctx, id, customer)
	return
}

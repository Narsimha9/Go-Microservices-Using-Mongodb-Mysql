package main

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   OrderService
}

func (mw loggingMiddleware) Create(ctx context.Context, order Order) (id string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Create",
			"CustomerID",
			order.CustomerID,
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return mw.next.Create(ctx, order)
}

//

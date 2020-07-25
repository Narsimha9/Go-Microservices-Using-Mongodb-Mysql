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
	next           OrderService
}

func (mw instrumentingMiddleware) Create(ctx context.Context, order Order) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "create", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output, err = mw.next.Create(ctx, order)
	return
}

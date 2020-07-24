package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type orderStruct struct {
	repository Repository
	logger     log.Logger
}

type OrderService interface {
	Create(ctx context.Context, order Order) (string, error)
}

var (
	ErrOrderNotFound   = errors.New("order not found")
	ErrCmdRepository   = errors.New("unable to command repository")
	ErrQueryRepository = errors.New("unable to query repository")
)

func NewService(rep Repository, logger log.Logger) OrderService {
	return &orderStruct{
		repository: rep,
		logger:     logger,
	}
}

func (o orderStruct) Create(ctx context.Context, order Order) (string, error) {
	fmt.Println("Crete method implementation called")
	logger := log.With(o.logger, "method", "Create")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	order.ID = id
	order.Status = "Pending"
	order.CreatedOn = time.Now().Unix()
	orders := Order{
		ID:           order.ID,
		CustomerID:   order.CustomerID,
		Status:       order.Status,
		CreatedOn:    order.CreatedOn,
		RestaurantId: order.RestaurantId,
		OrderItems:   order.OrderItems,
	}

	if err := o.repository.CreateOrder(ctx, orders); err != nil {
		fmt.Println("error generated")
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", id)

	return id, nil

}

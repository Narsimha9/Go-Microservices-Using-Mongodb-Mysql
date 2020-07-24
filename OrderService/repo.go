package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"gopkg.in/mgo.v2"
)

var RepoErr = errors.New("Unable to handle Repo Request")

const UserCollection = "Orders"

type repo struct {
	db     *mgo.Database
	logger log.Logger
}

func NewRepo(db *mgo.Database, logger log.Logger) (Repository, error) {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "mongodb"),
	}, nil
}

func (repo *repo) CreateOrder(ctx context.Context, order Order) error {
	fmt.Println("Create order in db called")
	err := db.C(UserCollection).Insert(order)
	if err != nil {
		fmt.Println("Error occured inside CreateOrder")
		return err
	} else {
		fmt.Println("Order Created:", err)
	}
	return nil
}

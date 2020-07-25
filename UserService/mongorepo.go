package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var RepoErr = errors.New("Unable to handle Repo Request")

const UserCollection = "gotestuser"

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

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	err := db.C(UserCollection).Insert(user)
	if err != nil {
		fmt.Println("Error occured inside CreateUser in repo")
		return err
	} else {
		fmt.Println("User Created:", user.Email)
	}
	return nil
}

func (repo *repo) GetUserById(ctx context.Context, id int) (interface{}, error) {
	coll := db.C(UserCollection)
	data := []User{}
	err := coll.Find(bson.M{"userid": id}).Select(bson.M{}).All(&data)
	if err != nil {
		fmt.Println("Error occured inside GetCUstomerById in repo")
		return "", err
	}
	return data, nil
}

func (repo *repo) GetAllUsers(ctx context.Context) (interface{}, error) {
	coll := db.C(UserCollection)
	data := []User{}
	err := coll.Find(bson.M{}).Select(bson.M{"id": 1, "userid": 1, "email": 1, "phone": 1}).All(&data)
	if err != nil {
		fmt.Println("Error occured inside GetCUstomerById in repo")
		return "", err
	}
	return data, nil
}

func (repo *repo) DeleteUser(ctx context.Context, id int) (string, error) {
	coll := db.C(UserCollection)
	err := coll.Remove(bson.M{"userid": id})
	if err != nil {
		fmt.Println("Error occured inside delete in repo")
		return "", err
	} else {
		msg := "user deleted successfully"
		return msg, nil
	}
}

func (repo *repo) UpdateUser(ctx context.Context, id int, user User) error {
	coll := db.C(UserCollection)
	err := coll.Update(bson.M{"userid": id}, bson.M{"$set": bson.M{"email": user.Email}})
	if err != nil {
		fmt.Println("Error occured inside update user repo")
		return err
	} else {
		return nil
	}

}

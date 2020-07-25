package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	// "reflect"
	"github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var RepoErr = errors.New("Unable to handle Repo Request")

const UserCollection = "gotestuser"

type repo struct {
	db     *mgo.Database
	logger log.Logger
}

type sqlrepo struct {
	sqldb  *sql.DB
	logger log.Logger
}

func NewRepo(db *mgo.Database, logger log.Logger) (Repository, error) {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "mongodb"),
	}, nil
}

func NewSqlRepo(sqldb *sql.DB, logger log.Logger) (Repository, error) {
	return &sqlrepo{
		sqldb:  sqldb,
		logger: log.With(logger, "sqlrepo", "sqldb"),
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

func (sqlrepo *sqlrepo) CreateUser(ctx context.Context, user User) error {
	fmt.Println("create user sql repo", sqldb)
	dbt := sqlrepo.sqldb
	sqlresp, err := dbt.Prepare("INSERT INTO gotestuser (id,userid,email,password,phone) VALUES (?,?,?,?,?)")
	sqlresp.Exec(user.Id, user.Userid, user.Email, user.Password, user.Phone)
	fmt.Println("User Created:", sqlresp, err)
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

func (sqlrepo *sqlrepo) GetUserById(ctx context.Context, id int) (interface{}, error) {
	user := User{}
	dbt := sqlrepo.sqldb
	err := dbt.QueryRow("SELECT * FROM gotestuser where userid = ?", id).Scan(&user.Id, &user.Userid, &user.Email, &user.Password, &user.Phone)

	if err != nil {
		fmt.Println("Error occured inside CreateUser in repo")
		return "", err
	}
	return user, nil
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

func (sqlrepo *sqlrepo) GetAllUsers(ctx context.Context) (interface{}, error) {
	data := []User{}
	fmt.Println("data is", data)
	dbt := sqlrepo.sqldb
	sqlresp, err := dbt.Query("select * from user")
	fmt.Println("in get All function", sqlresp, err)
	for sqlresp.Next() {
		var user User
		err = sqlresp.Scan(&user.Id, &user.Userid, &user.Email, &user.Password, &user.Phone)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("data", user)
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

func (sqlrepo *sqlrepo) DeleteUser(ctx context.Context, id int) (string, error) {
	dbt := sqlrepo.sqldb
	sqlresp, err := dbt.Query("DELETE FROM gotestuser where userid=?", id)
	fmt.Println("", sqlresp, err)

	return "", nil
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

func (sqlrepo *sqlrepo) UpdateUser(ctx context.Context, id int, user User) error {
	dbt := sqlrepo.sqldb
	sqlresp, err := dbt.Query("update gotestuser set email=? where userid=?", user.Email, id)
	fmt.Println("", sqlresp, err)
	return nil
}

// dbt:=sqlrepo.sqldb
// sqlresp, err := dbt.Query("select * from user")
// fmt.Println("in get All function",sqlresp,err)
// for sqlresp.Next() {
// 	data :=[]User{}
// 	s := reflect.ValueOf(&data).Elem()
// 	numCols := s.NumField()
// 	columns := make([]interface{}, numCols)
// 	for i := 0; i < numCols; i++ {
// 		field := s.Field(i)
// 		columns[i] = field.Addr().Interface()
// 	}
// 	err := rows.Scan(columns...)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(data)
// }
// return data, nil

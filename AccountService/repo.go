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

func (repo *repo) CreateCustomer(ctx context.Context, customer Customer) error {
	err := db.C(UserCollection).Insert(customer)
	if err != nil {
		fmt.Println("Error occured inside CreateCustomer in repo")
		return err
	} else {
		fmt.Println("User Created:", customer.Email)
	}
	return nil
}

func (sqlrepo *sqlrepo) CreateCustomer(ctx context.Context, customer Customer) error {
	fmt.Println("create customer sql repo", sqldb)
	dbt := sqlrepo.sqldb
	sqlresp, err := dbt.Prepare("INSERT INTO gotestuser (id,customerid,email,password,phone) VALUES (?,?,?,?,?)")
	sqlresp.Exec(customer.Id, customer.Customerid, customer.Email, customer.Password, customer.Phone)
	fmt.Println("User Created:", sqlresp, err)
	return nil
}

func (repo *repo) GetCustomerById(ctx context.Context, id int) (interface{}, error) {
	coll := db.C(UserCollection)
	data := []Customer{}
	err := coll.Find(bson.M{"customerid": id}).Select(bson.M{}).All(&data)
	if err != nil {
		fmt.Println("Error occured inside GetCUstomerById in repo")
		return "", err
	}
	return data, nil
}

func (sqlrepo *sqlrepo) GetCustomerById(ctx context.Context, id int) (interface{}, error) {
	customer := Customer{}
	dbt := sqlrepo.sqldb
	err := dbt.QueryRow("SELECT * FROM gotestuser where customerid = ?", id).Scan(&customer.Id, &customer.Customerid, &customer.Email, &customer.Password, &customer.Phone)

	if err != nil {
		fmt.Println("Error occured inside CreateCustomer in repo")
		return "", err
	}
	return customer, nil
}

func (repo *repo) GetAllCustomers(ctx context.Context) (interface{}, error) {
	coll := db.C(UserCollection)
	data := []Customer{}
	err := coll.Find(bson.M{}).Select(bson.M{"id": 1, "customerid": 1, "email": 1, "phone": 1}).All(&data)
	if err != nil {
		fmt.Println("Error occured inside GetCUstomerById in repo")
		return "", err
	}
	return data, nil
}

func (sqlrepo *sqlrepo) GetAllCustomers(ctx context.Context) (interface{}, error) {
	data := []Customer{}
	fmt.Println("data is", data)
	dbt := sqlrepo.sqldb
	sqlresp, err := dbt.Query("select * from user")
	fmt.Println("in get All function", sqlresp, err)
	for sqlresp.Next() {
		var customer Customer
		err = sqlresp.Scan(&customer.Id, &customer.Customerid, &customer.Email, &customer.Password, &customer.Phone)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("data", customer)
	}
	return data, nil
}

func (repo *repo) DeleteCustomer(ctx context.Context, id int) (string, error) {
	coll := db.C(UserCollection)
	err := coll.Remove(bson.M{"customerid": id})
	if err != nil {
		fmt.Println("Error occured inside delete in repo")
		return "", err
	} else {
		msg := "customer deleted successfully"
		return msg, nil
	}
}

func (sqlrepo *sqlrepo) DeleteCustomer(ctx context.Context, id int) (string, error) {
	dbt := sqlrepo.sqldb
	sqlresp, err := dbt.Query("DELETE FROM gotestuser where customerid=?", id)
	fmt.Println("", sqlresp, err)

	return "", nil
}

func (repo *repo) UpdateCustomer(ctx context.Context, id int, customer Customer) error {
	coll := db.C(UserCollection)
	err := coll.Update(bson.M{"customerid": id}, bson.M{"$set": bson.M{"email": customer.Email}})
	if err != nil {
		fmt.Println("Error occured inside update customer repo")
		return err
	} else {
		return nil
	}

}

func (sqlrepo *sqlrepo) UpdateCustomer(ctx context.Context, id int, customer Customer) error {
	dbt := sqlrepo.sqldb
	sqlresp, err := dbt.Query("update gotestuser set email=? where customerid=?", customer.Email, id)
	fmt.Println("", sqlresp, err)
	return nil
}

// dbt:=sqlrepo.sqldb
// sqlresp, err := dbt.Query("select * from user")
// fmt.Println("in get All function",sqlresp,err)
// for sqlresp.Next() {
// 	data :=[]Customer{}
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

package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
)

var MysqlRepoErr = errors.New("Unable to handle Repo Request")

const UsersCollection = "user"

type sqlrepo struct {
	sqldb  *sql.DB
	logger log.Logger
}

func NewSqlRepo(sqldb *sql.DB, logger log.Logger) (Repository, error) {
	return &sqlrepo{
		sqldb:  sqldb,
		logger: log.With(logger, "sqlrepo", "sqldb"),
	}, nil
}

func (sqlrepo *sqlrepo) CreateUser(ctx context.Context, user User) error {
	fmt.Println("create user sql repo", sqldb)
	dbt := sqlrepo.sqldb
	sqlresp, err := dbt.Prepare("INSERT INTO user (id,userid,email,password,phone) VALUES (?,?,?,?,?)")
	sqlresp.Exec(user.Id, user.Userid, user.Email, user.Password, user.Phone)
	fmt.Println("User Created:", sqlresp, err)
	return nil
}

func (sqlrepo *sqlrepo) GetUserById(ctx context.Context, id int) (interface{}, error) {
	user := User{}
	dbt := sqlrepo.sqldb
	err := dbt.QueryRow("SELECT * FROM user where userid = ?", id).Scan(&user.Id, &user.Userid, &user.Email, &user.Password, &user.Phone)

	if err != nil {
		fmt.Println("Error occured inside CreateUser in repo")
		return "", err
	}
	return user, nil
}

func (sqlrepo *sqlrepo) GetAllUsers(ctx context.Context) (interface{}, error) {
	fmt.Println("Getall users in repo")
	user := User{}
	var res []interface{}
	results, err := sqlrepo.sqldb.QueryContext(ctx, "SELECT user.userid,user.email,user.phone FROM user")
	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}
	defer results.Close()
	for results.Next() {
		err = results.Scan(&user.Userid, &user.Email, &user.Phone)
		res = append([]interface{}{user}, res...)
	}
	return res, nil
}

func (sqlrepo *sqlrepo) DeleteUser(ctx context.Context, id int) (string, error) {
	dbt := sqlrepo.sqldb
	sqlresp, err := dbt.Query("DELETE FROM user where userid=?", id)
	fmt.Println("", sqlresp, err)

	return "", nil
}

func (sqlrepo *sqlrepo) UpdateUser(ctx context.Context, id int, user User) error {
	dbt := sqlrepo.sqldb
	sqlresp, err := dbt.Query("update user set email=? where userid=?", user.Email, id)
	fmt.Println("", sqlresp, err)
	return nil
}

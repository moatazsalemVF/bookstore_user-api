package mysql

import (
	"database/sql"
	"fmt"
	"os"

	//import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsername = "os_env_mysql_users_username"
	mysqlPassword = "os_env_mysql_users_password"
	mysqlHost     = "os_env_mysql_users_host"
	mysqlSchema   = "os_env_mysql_users_schema"
)

var (
	//Client is a DB connection to users DB
	Client *sql.DB

	username = os.Getenv("os_env_mysql_users_username")
	password = os.Getenv("os_env_mysql_users_password")
	host     = os.Getenv("os_env_mysql_users_host")
	schema   = os.Getenv("os_env_mysql_users_schema")
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	Client = db
}

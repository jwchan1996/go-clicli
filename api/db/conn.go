package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
//root:changhao123@tcp(localhost:3306)/acgzone?charset=utf8
//root:root@tcp(localhost:3307)/uraban?charset=utf8
	dbConn, err = sql.Open("mysql", "root:changhao123@tcp(localhost:3306)/acgzone?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}

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
//root:changhao2333@tcp(localhost:3306)/acgzone?charset=utf8
//root:root@tcp(localhost:3306)/acgzone?charset=utf8
	dbConn, err = sql.Open("mysql", "root:changhao2333@tcp(localhost:3306)/clicli?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}

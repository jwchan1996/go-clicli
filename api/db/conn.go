package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DbConn *sql.DB
	err    error
)

func init() {
	// DbConn, err = sql.Open("mysql", "root:changhao2333@tcp(api.clicli.us:3306)/clicli?charset=utf8")
	DbConn, err = sql.Open("mysql", "ppap:ppap@tcp(localhost:3306)/clicli?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	log.Println("ok")
}

//root:root@tcp(localhost:3306)/clicli?charset=utf8

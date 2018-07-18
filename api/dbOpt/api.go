package dbOpt

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func openConn() *sql.DB{

	return dbConn
}

func CreateUserCredential(name:string, pwd:string) error {

}

func GetUserCredential(name:string) (string,error) {

}
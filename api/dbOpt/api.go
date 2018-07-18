package dbOpt

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func CreateUserCredential(name string, pwd string, role string, qq int) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (name,pwd,role,qq) VALUES (?,?,?,?)")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	_, err = stmtIns.Exec(name, pwd, role, qq)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(name string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(name).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()

	return pwd, nil
}

func DeleteUser(name string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE name=? AND pwd=?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmtDel.Exec(name, pwd)
	if err != nil {
		return err
	}
	stmtDel.Close()

	return nil
}

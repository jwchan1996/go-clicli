package db

import (
	"log"
	"github.com/132yse/acgzone-server/api/util"
	"github.com/132yse/acgzone-server/api/def"
	"database/sql"
)

func CreateUser(name string, pwd string, role string, qq int, sign string) error {
	pwd = util.Cipher(pwd)
	stmtIns, err := dbConn.Prepare("INSERT INTO users (name,pwd,role,qq,sign) VALUES (?,?,?,?,?)")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	_, err = stmtIns.Exec(name, pwd, role, qq, sign)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func UpdateUser(id int, name string, pwd string, role string, qq int, sign string) (*def.UserCredential, error) {
	stmtIns, err := dbConn.Prepare("UPDATE users SET name=?,pwd=?,role=?,qq=?,sign=? WHERE id =?")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(&name, &pwd, &role, &qq, &sign, &id)
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()

	res := &def.UserCredential{Id: id, Name: name, Pwd: pwd, QQ: qq, Role: role, Desc: sign}
	defer stmtIns.Close()
	return res, err
}

func GetUser(name string) (*def.UserCredential, error) {
	stmtOut, err := dbConn.Prepare("SELECT id,pwd,role,qq,sign FROM users WHERE name = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var id, qq int
	var pwd, role, sign string
	err = stmtOut.QueryRow(name).Scan(&id, &pwd, &role, &qq, &sign)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &def.UserCredential{Id: id, Pwd: pwd, Name: name, Role: role, QQ: qq, Desc: sign}

	defer stmtOut.Close()

	return res, nil
}

func GetUsers(role string, page int, pageSize int) ([]*def.UserCredential, error) {
	start := pageSize * (page - 1)

	stmtOut, err := dbConn.Prepare("SELECT id, name, role, qq, sign FROM users WHERE role =? limit ?,?")

	var res []*def.UserCredential

	rows, err := stmtOut.Query(role, start, pageSize)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, qq int
		var name, role, sign string
		if err := rows.Scan(&id, &name, &role, &qq, &sign); err != nil {
			return res, err
		}

		c := &def.UserCredential{Id: id, Name: name, Role: role, QQ: qq, Desc: sign}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil

}

func SearchUsers(key string) ([]*def.UserCredential, error) {
	key = string("%" + key + "%")
	stmtOut, err := dbConn.Prepare("SELECT id, name, role, qq, sign FROM users WHERE name LIKE ?")

	var res []*def.UserCredential

	rows, err := stmtOut.Query(key)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, qq int
		var name, role, sign string
		if err := rows.Scan(&id, &name, &role, &qq, &sign); err != nil {
			return res, err
		}

		c := &def.UserCredential{Id: id, Name: name, Role: role, QQ: qq, Desc: sign}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil

}

func DeleteUser(id int) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE id =?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmtDel.Exec(id)
	if err != nil {
		return err
	}
	stmtDel.Close()

	return nil
}

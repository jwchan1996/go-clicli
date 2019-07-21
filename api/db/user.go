package db

import (
	"log"
	"github.com/132yse/acgzone-server/api/util"
	"github.com/132yse/acgzone-server/api/def"
	"fmt"
	"database/sql"
)

func CreateUser(name string, pwd string, level int, qq string, sign string) error {
	pwd = util.Cipher(pwd)
	stmtIns, err := dbConn.Prepare("INSERT INTO users (name,pwd,level,qq,sign) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(name, pwd, level, qq, sign)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func UpdateUser(id int, name string, pwd string, level int, qq string, sign string) (*def.User, error) {
	if pwd == "" {
		stmtIns, err := dbConn.Prepare("UPDATE users SET name=?,level=?,qq=?,sign=? WHERE id =?")
		if err != nil {
			return nil, err
		}
		_, err = stmtIns.Exec(&name, &level, &qq, &sign, &id)
		if err != nil {
			return nil, err
		}
		defer stmtIns.Close()

		res := &def.User{Id: id, Name: name, QQ: qq, Level: level, Desc: sign}
		defer stmtIns.Close()
		return res, err
	} else {
		pwd = util.Cipher(pwd)
		stmtIns, err := dbConn.Prepare("UPDATE users SET name=?,pwd=?,level=?,qq=?,sign=? WHERE id =?")
		if err != nil {
			return nil, err
		}
		_, err = stmtIns.Exec(&name, &pwd, &level, &qq, &sign, &id)
		if err != nil {
			return nil, err
		}
		defer stmtIns.Close()

		res := &def.User{Id: id, Name: name, Pwd: pwd, QQ: qq, Level: level, Desc: sign}
		defer stmtIns.Close()
		return res, err
	}

}

func GetUser(name string, id int) (*def.User, error) {
	var query string
	if name != "" {
		query += `SELECT id,name,pwd,level,qq,sign FROM users WHERE name = ?`
	} else {
		query += `SELECT id,name,pwd,level,qq,sign FROM users WHERE id = ?`
	}
	stmt, _ := dbConn.Prepare(query)
	var level int
	var pwd, sign, qq string
	if name != "" {
		err = stmt.QueryRow(name).Scan(&id, &name, &pwd, &level, &qq, &sign)
	} else {
		err = stmt.QueryRow(id).Scan(&id, &name, &pwd, &level, &qq, &sign)
	}

	defer stmt.Close()

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &def.User{Id: id, Pwd: pwd, Name: name, Level: level, QQ: qq, Desc: sign}

	return res, nil
}

func GetUsers(level int, page int, pageSize int) ([]*def.User, error) {
	start := pageSize * (page - 1)

	var query string
	if level > 4 {
		query += `NOT level = 'user'`
	} else if level > 0 {
		query += fmt.Sprintf(`level = '%s'`, string(level))
	}
	rawSql := fmt.Sprintf(`SELECT id, name, level, qq, sign FROM users WHERE %s limit ?,?`, query)
	stmtOut, err := dbConn.Prepare(rawSql)

	var res []*def.User

	rows, err := stmtOut.Query(start, pageSize)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, level int
		var name, sign, qq string
		if err := rows.Scan(&id, &name, &level, &qq, &sign); err != nil {
			return res, err
		}

		c := &def.User{Id: id, Name: name, Level: level, QQ: qq, Desc: sign}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil

}

func SearchUsers(key string) ([]*def.User, error) {
	key = string("%" + key + "%")
	stmtOut, err := dbConn.Prepare("SELECT id, name, level, qq, sign FROM users WHERE name LIKE ?")

	var res []*def.User

	rows, err := stmtOut.Query(key)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, level int
		var name, sign, qq string
		if err := rows.Scan(&id, &name, &level, &qq, &sign); err != nil {
			return res, err
		}

		c := &def.User{Id: id, Name: name, Level: level, QQ: qq, Desc: sign}
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

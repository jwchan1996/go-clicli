package db

import (
	"log"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/132yse/acgzone-server/api/util"
)

//用户增删改查

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

//文章增删改查

func AddPost(title string, content string, status string, sort string) (*def.Post, error) {
	t := time.Now()
	ctime := t.Format("2006-01-02")
	stmtIns, err := dbConn.Prepare("INSERT INTO posts (title,content,status,sort,time) VALUES (?,?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(title, content, status, sort, ctime)
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()

	res := &def.Post{Title: title, Content: content, Status: status, Sort: sort, Time: ctime}
	defer stmtIns.Close()
	return res, err
}

func GetPost(id int) (*def.Post, error) {
	stmtOut, err := dbConn.Prepare("SELECT id,title,content,status,sort,time FROM posts WHERE id = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	var pid int
	var title, content, status, sort, ctime string

	err = stmtOut.QueryRow(id).Scan(&pid, &title, &content, &status, &sort, &ctime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &def.Post{Id: pid, Title: title, Content: content, Status: status, Sort: sort, Time: time}

	return res, nil
}

//查找发布状态的所有文章

func GetPostsByStatus(status string) ([]*def.Post, error) {
	stmtOut, err := dbConn.Prepare("SELECT id, title, content, status, sort, time FROM posts WHERE status =?")

	var res []*def.Post

	rows, err := stmtOut.Query(status)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id int
		var title, content, status, sort, ctime string
		if err := rows.Scan(&id, &title, &content, &status, &sort, &ctime); err != nil {
			return res, err
		}

		c := &def.Post{Id: id, Title: title, Content: content, Status: status, Sort: sort, Time: ctime}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil

}

func DeletePost(id int) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM posts WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(id)
	if err != nil {
		return err
	}
	stmtDel.Close()

	return nil

}

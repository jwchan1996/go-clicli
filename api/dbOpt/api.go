package dbOpt

import (
	"time"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/132yse/acgzone-server/api/def"
)

//用户增删改查

func CreateUser(name string, pwd string, role string, qq int) error {
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

func GetUser(name string) (string, error) {
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

//文章增删改查

func AddPost(title string, content string, sort string, status string) (*def.Post, error) {
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
	stmtOut, err := dbConn.Prepare("SELECT title,content,status,sort,time FROM posts WHERE id = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var title string
	var content string
	var status string
	var sort string
	var time string

	err = stmtOut.QueryRow(id).Scan(&title, &content, &status, &sort, &time)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &def.Post{Title: title, Content: content, Status: status, Sort: sort, Time: time}

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

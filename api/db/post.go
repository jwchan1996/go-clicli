package db

import (
	"github.com/132yse/acgzone-server/api/def"
	"time"
	"database/sql"
	"log"
)

func AddPost(title string, content string, status string, sort string, tag string, uid int) (*def.Post, error) {
	t := time.Now()
	ctime := t.Format("2006-01-02 15:04")
	stmtIns, err := dbConn.Prepare("INSERT INTO posts (title,content,status,sort,type,time,uid) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		return nil, err

	}
	_, err = stmtIns.Exec(title, content, status, sort, tag, ctime, uid)
	if err != nil {

		return nil, err
	}
	defer stmtIns.Close()

	res := &def.Post{Title: title, Content: content, Status: status, Sort: sort, Type: tag, Time: ctime, Uid: uid}
	defer stmtIns.Close()

	return res, err
}

func UpdatePost(id int, title string, content string, status string, sort string, tag string) (*def.Post, error) {
	t := time.Now()
	ctime := t.Format("2006-01-02 15:04")
	stmtIns, err := dbConn.Prepare("UPDATE posts SET title=?,content=?,status=?,sort=?,type=?,time=? WHERE id =?")
	if err != nil {
		return nil, err

	}
	_, err = stmtIns.Exec(&title, &content, &status, &sort, &ctime, &id, &tag)
	if err != nil {

		return nil, err
	}
	defer stmtIns.Close()

	res := &def.Post{Id: id, Title: title, Content: content, Status: status, Sort: sort, Time: ctime}
	defer stmtIns.Close()
	return res, err
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

func GetPost(id int) (*def.Post, error) {
	stmtOut, err := dbConn.Prepare(`SELECT posts.id,posts.title,posts.content,posts.status,posts.sort,posts.type,posts.time,users.id,users.name,users.qq FROM posts 
INNER JOIN users ON posts.uid = users.id WHERE posts.id = ?`)
	if err != nil {
		return nil, err
	}
	var pid, uid int
	var title, content, status, sort, tag, ctime, uname, uqq string

	err = stmtOut.QueryRow(id).Scan(&pid, &title, &content, &status, &sort, &tag, &ctime, &uid, &uname, &uqq)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &def.Post{Id: pid, Title: title, Content: content, Status: status, Sort: sort, Type: tag, Time: ctime, Uid: uid, Uname: uname, Uqq: uqq}

	return res, nil
}

func GetPostsOneOf(status string, sort string, tag string, uid int, page int, pageSize int) ([]*def.Post, error) {
	start := pageSize * (page - 1)

	stmtOut, err := dbConn.Prepare(`SELECT posts.id,posts.title,posts.content,posts.status,posts.sort,posts.type,posts.time,users.id,users.name,users.qq FROM posts INNER JOIN users ON posts.uid = users.id 
WHERE posts.status =? OR posts.sort=? OR posts.type=? OR posts.uid =? ORDER BY time DESC limit ?,?`)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var res []*def.Post

	rows, err := stmtOut.Query(status, sort, tag, uid, start, pageSize)
	if err != nil {
		log.Printf("%s", err)
		return res, err
	}

	for rows.Next() {
		var id, uid int
		var title, content, status, sort, tag, ctime, uname, uqq string
		if err := rows.Scan(&id, &title, &content, &status, &sort, &tag, &ctime, &uid, &uname, &uqq); err != nil {
			log.Printf("%s", err)
			return res, err
		}

		c := &def.Post{Id: id, Title: title, Content: content, Status: status, Sort: sort, Type: tag, Time: ctime, Uid: uid, Uname: uname, Uqq: uqq}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil

}

func GetPostsBoth(status string, sort string, tag string, uid int, page int, pageSize int) ([]*def.Post, error) {
	start := pageSize * (page - 1)
	log.Printf("%x", start)

	stmtOut, err := dbConn.Prepare(`SELECT posts.id,posts.title,posts.content,posts.status,posts.sort,posts.type,posts.time,users.id,users.name,users.qq FROM posts LEFT JOIN users ON posts.uid = users.id 
WHERE (posts.sort=? OR posts.uid =? OR posts.type=?) AND posts.status =? ORDER BY time DESC limit ?,?`)

	if err != nil {
		return nil, err
	}

	var res []*def.Post

	rows, err := stmtOut.Query(sort, uid, tag, status, start, pageSize)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, uid int
		var title, content, status, sort, tag, ctime, uname, uqq string
		if err := rows.Scan(&id, &title, &content, &status, &sort, &tag, &ctime, &uid, &uname, &uqq); err != nil {
			return res, err
		}

		c := &def.Post{Id: id, Title: title, Content: content, Status: status, Sort: sort, Type: tag, Time: ctime, Uid: uid, Uname: uname, Uqq: uqq}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil

}

func SearchPosts(key string) ([]*def.Post, error) {
	key = string("%" + key + "%")
	stmtOut, err := dbConn.Prepare("SELECT posts.id, posts.title, posts.content, posts.status, posts.sort, posts.type,posts.time,users.id,users.name,users.qq FROM posts LEFT JOIN users ON posts.uid = users.id WHERE title LIKE ? OR content LIKE ?")

	var res []*def.Post

	rows, err := stmtOut.Query(key, key)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, uid int
		var title, content, status, sort, tag, ctime, uname, uqq string
		if err := rows.Scan(&id, &title, &content, &status, &sort, &tag, &ctime, &uid, &uname, &uqq); err != nil {
			return res, err
		}

		c := &def.Post{Id: id, Title: title, Content: content, Status: status, Sort: sort, Type: tag, Time: ctime, Uid: uid, Uname: uname, Uqq: uqq}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil
}

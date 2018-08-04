package db

import (
	"github.com/132yse/acgzone-server/api/def"
	"time"
	"log"
)

func AddComment(content string, pid int, uid int) (*def.Comment, error) {
	t := time.Now()
	ctime := t.Format("2006-01-02 15:04")
	stmtIns, err := dbConn.Prepare("INSERT INTO comments (content,time,pid,uid) VALUES (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(content, ctime, pid, uid)
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()

	res := &def.Comment{Content: content, Time: ctime, Uid: uid, Pid: pid}
	defer stmtIns.Close()
	return res, err
}

func GetComments(pid int, uid int, page int, pageSize int) ([]*def.Comment, error) {
	start := pageSize * (page - 1)

	stmtOut, err := dbConn.Prepare(`SELECT comments.id,comments.content,comments.time,comments.pid,users.id,users.name,users.qq FROM comments INNER JOIN users ON comments.uid = users.id 
WHERE comments.pid=? OR comments.uid =? ORDER BY time DESC limit ?,?`)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var res []*def.Comment

	rows, err := stmtOut.Query(pid, uid, start, pageSize)
	if err != nil {
		log.Printf("%s", err)
		return res, err
	}

	for rows.Next() {
		var id, pid, uid int
		var content, ctime, uname, uqq string
		if err := rows.Scan(&id, &content, &ctime, &pid, &uid, &uname, &uqq); err != nil {
			log.Printf("%s", err)
			return res, err
		}

		c := &def.Comment{Id: id, Content: content, Time: ctime, Pid: pid, Uid: uid, Uname: uname, Uqq: uqq}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil

}

func DeleteComment(id int) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM comments WHERE id=?")
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

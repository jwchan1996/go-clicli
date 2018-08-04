package db

import (
	"github.com/132yse/acgzone-server/api/def"
	"time"
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

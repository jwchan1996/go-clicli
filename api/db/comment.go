package db

import (
	"github.com/132yse/acgzone-server/api/def"
	"log"
	"time"
)

func AddComment(content string, pid int, uid int, tuid int, vid int, dtime int, color string) (*def.Comment, error) {
	t := time.Now()
	ctime := t.Format("2006-01-02 15:04")
	stmtIns, err := dbConn.Prepare("INSERT INTO comments (content,ctime,pid,uid,tuid,vid,time,color) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(content, ctime, pid, uid, tuid, vid, dtime, color)
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()

	res := &def.Comment{Content: content, Ctime: ctime, Uid: uid, Pid: pid, Tuid: tuid, Vid: vid, Time: dtime, Color: color}
	defer stmtIns.Close()
	return res, err
}

func GetComments(pid int, uid int, vid int, page int, pageSize int) ([]*def.Comment, error) {
	start := pageSize * (page - 1)

	stmtOut, err := dbConn.Prepare(`SELECT comments.id,comments.content,comments.ctime,comments.pid,comments.vid,comments.tuid,comments.time,comments.color,users.id,users.name,users.qq FROM comments INNER JOIN users ON comments.uid = users.id 
WHERE comments.pid=? OR comments.uid =? OR comments.vid =? ORDER BY ctime DESC limit ?,?`)

	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var res []*def.Comment

	rows, err := stmtOut.Query(pid, uid, vid, start, pageSize)
	if err != nil {
		log.Printf("%s", err)
		return res, err
	}

	for rows.Next() {
		var id, pid, uid, vid, tuid, dtime int
		var content, ctime, uname, uqq, color string
		if err := rows.Scan(&id, &content, &ctime, &pid, &vid, &tuid, &dtime, &color, &uid, &uname, &uqq); err != nil {
			log.Printf("%s", err)
			return res, err
		}

		c := &def.Comment{Id: id, Content: content, Ctime: ctime, Pid: pid, Vid: vid, Tuid: tuid, Time: dtime, Color: color, Uid: uid, Uname: uname, Uqq: uqq}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil

}

func DeleteComment(id int, pid int) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM comments WHERE id=? OR pid=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(id, pid)
	if err != nil {
		return err
	}
	stmtDel.Close()

	return nil

}

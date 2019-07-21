package db

import (
	"database/sql"
	"github.com/132yse/acgzone-server/api/def"
	"time"
)

func AddVideo(oid int, title string, content string, pid int, uid int) (*def.Video, error) {
	t := time.Now()
	ctime := t.Format("2006-01-02 15:04")
	stmtIns, err := dbConn.Prepare("INSERT INTO videos (oid,title,content,time,pid,uid) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(oid, title, content, ctime, pid, uid)
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()

	res := &def.Video{Oid: oid, Title: title, Content: content, Time: ctime, Uid: uid, Pid: pid}
	defer stmtIns.Close()
	return res, err
}

func GetVideos(pid int, uid int, page int, pageSize int) ([]*def.Video, error) {
	start := pageSize * (page - 1)

	stmtOut, err := dbConn.Prepare(`SELECT videos.id,videos.oid,videos.title,videos.content,videos.time,videos.pid,users.id,users.name,users.qq FROM videos INNER JOIN users ON videos.uid = users.id 
WHERE videos.pid=? OR videos.uid =? ORDER BY oid limit ?,?`)

	if err != nil {
		return nil, err
	}

	var res []*def.Video

	rows, err := stmtOut.Query(pid, uid, start, pageSize)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, oid, pid, uid int
		var title, content, ctime, uname, uqq string
		if err := rows.Scan(&id, &oid, &title, &content, &ctime, &pid, &uid, &uname, &uqq); err != nil {
			return res, err
		}

		c := &def.Video{Id: id, Oid: oid, Title: title, Content: content, Time: ctime, Pid: pid, Uid: uid, Uname: uname, Uqq: uqq}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil

}

func GetVideo(id int) (*def.Video, error) {
	stmtOut, err := dbConn.Prepare(`SELECT videos.id,videos.oid,videos.title,videos.content,videos.time,posts.id,posts.title,users.id,users.name,users.qq FROM (videos INNER JOIN posts ON videos.pid=posts.id) 
INNER JOIN users ON videos.uid = users.id WHERE videos.id = ?`)
	if err != nil {
		return nil, err
	}
	var vid, uid, oid, pid int
	var title, content, ctime, uname, uqq, ptitle string

	err = stmtOut.QueryRow(id).Scan(&vid, &oid, &title, &content, &ctime, &pid, &ptitle, &uid, &uname, &uqq)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmtOut.Close()

	res := &def.Video{Id: vid, Oid: oid, Title: title, Content: content, Time: ctime, Pid: pid, Ptitle: ptitle, Uid: uid, Uname: uname, Uqq: uqq}

	return res, nil
}

func UpdateVideo(id int, oid int, title string, content string, pid int, uid int) (*def.Video, error) {
	t := time.Now()
	ctime := t.Format("2006-01-02 15:04")
	stmtIns, err := dbConn.Prepare("UPDATE videos SET oid=?,title=?,content=?,pid=?,uid=?,time=? WHERE id =?")
	if err != nil {
		return nil, err

	}
	_, err = stmtIns.Exec(&oid, &title, &content, &pid, &uid, &ctime, &id)
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()

	res := &def.Video{Id: id, Oid: oid, Title: title, Content: content, Pid: pid, Uid: uid, Time: ctime}
	defer stmtIns.Close()
	return res, err
}

func DeleteVideo(id int, pid int) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM videos WHERE id=? OR pid=?")
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

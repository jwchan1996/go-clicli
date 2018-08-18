package db

import (
	"github.com/132yse/acgzone-server/api/def"
	"log"
	"database/sql"
)

func AddPageView(pv int, pid int) error {

	stmtIns, err := dbConn.Prepare("UPDATE pv SET pv=? WHERE pid =?")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(&pv, &pid)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetCount(pid int) (*def.Count, error) {
	stmtOut, err := dbConn.Prepare("SELECT pid,pv FROM pv WHERE pid = ?")
	stmtCount, err := dbConn.Prepare("SELECT COUNT(*) FROM comments WHERE pid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var pv, cv int
	err = stmtOut.QueryRow(pid).Scan(&pid, &pv)
	err = stmtCount.QueryRow(pid).Scan(&pid, &cv)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &def.Count{Pid: pid, Pv: pv, Cv: cv}

	defer stmtOut.Close()

	return res, nil
}

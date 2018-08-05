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

func GetPageView(pid int) (*def.Pv, error) {
	stmtOut, err := dbConn.Prepare("SELECT pid,pv FROM pv WHERE pid = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var pv int
	err = stmtOut.QueryRow(pid).Scan(&pid, &pv)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &def.Pv{Pid: pid, Pv: pv}

	defer stmtOut.Close()

	return res, nil
}

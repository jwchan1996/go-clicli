package db

import (
	"github.com/cliclitv/go-clicli/def"
	"database/sql"
)

func GetPv(pid int) (*def.Pv, error) {
	stmtCount, err := dbConn.Prepare("SELECT pv FROM pv WHERE pid = ?")
	if err != nil {
		return nil, err
	}

	var pv int
	err = stmtCount.QueryRow(pid).Scan(&pv)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	res := &def.Pv{Pid: pid, Pv: pv}

	defer stmtCount.Close()

	return res, nil
}

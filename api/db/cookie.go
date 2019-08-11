package db

import (
	"database/sql"
	"github.com/132yse/acgzone-server/api/def"
	"log"
)

func ReplaceCookie(uid int, hcy string, quqi string) (*def.Cookie, error) {
	stmtIns, err := dbConn.Prepare("REPLACE INTO cookies (uid,hcy,quqi) VALUES (?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(uid, hcy, quqi)
	if err != nil {
		return nil, err
	}
	defer stmtIns.Close()

	res := &def.Cookie{Uid: uid, Hcy: hcy, Quqi: quqi}
	defer stmtIns.Close()
	return res, err
}

func GetCookie(uid int) (*def.Cookie, error) {
	stmtOut, err := dbConn.Prepare("SELECT uid,hcy,quqi FROM cookies WHERE uid = ?")
	if err != nil {
		return nil, err
	}

	var hcy, quqi string
	err = stmtOut.QueryRow(uid).Scan(&uid, &hcy, &quqi)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%v", err)
		return nil, err
	}
	res := &def.Cookie{Uid: uid, Hcy: hcy, Quqi: quqi}

	defer stmtOut.Close()

	return res, nil
}

package handler

import (
	"encoding/json"
	"github.com/132yse/acgzone-server/api/db"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func ReplaceCookie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	cbody := &def.Cookie{}

	if err := json.Unmarshal(req, cbody); err != nil {
		sendMsg(w, 401, "参数解析失败")
		log.Printf("%v", err)
		return
	}

	if resp, err := db.ReplaceCookie(cbody.Uid, cbody.Hcy); err != nil {
		sendMsg(w, 401, "数据库错误")
		return
	} else {
		res := def.Cookie{Uid: resp.Uid, Hcy: resp.Hcy}
		sendCookieResponse(w, res, 200)
	}

}

func GetCookie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid, _ := strconv.Atoi(p.ByName("uid"))
	resp, err := db.GetCookie(uid)
	if err != nil || resp == nil {
		sendMsg(w, 401, "数据库错误")
		log.Printf("%v", err)
		return
	}
	res := def.Cookie{Uid: resp.Uid, Hcy: resp.Hcy, Tyyp: resp.Tyyp, Bit: resp.Bit}
	sendCookieResponse(w, res, 200)

}

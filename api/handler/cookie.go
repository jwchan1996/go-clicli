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
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		log.Printf("%v", err)
		return
	}

	if resp, err := db.ReplaceCookie(cbody.Uid, cbody.Hcy, cbody.Tyyp, cbody.Bit); err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Cookie{Uid: resp.Uid, Hcy: resp.Hcy, Tyyp: resp.Tyyp, Bit: resp.Bit}
		sendCookieResponse(w, res, 201)
	}

}

func GetCookie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid, _ := strconv.Atoi(p.ByName("uid"))
	resp, err := db.GetCookie(uid)
	if err != nil {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		log.Printf("%v", err)
		return
	}
	if resp == nil{
		sendErrorResponse(w, def.ErrorDB)
		return
	}
	res := def.Cookie{Uid: resp.Uid, Hcy: resp.Hcy, Tyyp: resp.Tyyp, Bit: resp.Bit}
	sendCookieResponse(w, res, 201)

}

package handler

import (
	"encoding/json"
	"github.com/cliclitv/go-clicli/db"
	"github.com/cliclitv/go-clicli/def"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func ReplacePv(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	cbody := &def.Pv{}

	if err := json.Unmarshal(req, cbody); err != nil {
		sendMsg(w, 401, "参数解析失败")
		return
	}

	if resp, err := db.ReplacePv(cbody.Pid, cbody.Pv); err != nil {
		sendMsg(w, 401, "数据库错误")
		return
	} else {
		res := def.Pv{Pid: resp.Pid, Pv: resp.Pv}
		sendPvResponse(w, res, 200)
	}

}

func GetPv(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid, _ := strconv.Atoi(p.ByName("pid"))
	resp, err := db.GetPv(uid)
	if err != nil || resp == nil {
		sendMsg(w, 401, "数据库错误")
		log.Printf("%v", err)
		return
	}
	res := def.Pv{Pid: resp.Pid, Pv: resp.Pv}
	sendPvResponse(w, res, 200)

}
package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"github.com/132yse/acgzone-server/api/db"
	"github.com/132yse/acgzone-server/api/def"
)

func GetCommentCount(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Cross(w)
	pid, _ := strconv.Atoi(p.ByName("pid"))
	resp, err := db.GetCommentCount(pid)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Count{Pid: resp.Pid, Pv: resp.Pv, Cv: resp.Cv}
		sendCountResponse(w, res, 201)
	}
}

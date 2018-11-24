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

func AddComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	req, _ := ioutil.ReadAll(r.Body)
	cbody := &def.Comment{}

	if err := json.Unmarshal(req, cbody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if resp, err := db.AddComment(cbody.Content, cbody.Pid, cbody.Uid, cbody.Tuid, cbody.Vid, cbody.Time, cbody.Color); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Comment{Id: resp.Id, Content: resp.Content, Time: resp.Time, Pid: resp.Pid, Uid: resp.Uid}
		sendCommentResponse(w, res, 201)
	}

}

func GetComments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pid, _ := strconv.Atoi(r.URL.Query().Get("pid"))
	uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	vid, _ := strconv.Atoi(r.URL.Query().Get("vid"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	resp, err := db.GetComments(pid, uid, vid, page, pageSize)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		log.Printf("%s", err)
		return
	} else {
		res := &def.Comments{Comments: resp}
		sendCommentsResponse(w, res, 201)
	}
}

func DeleteComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	pid, _ := strconv.Atoi(r.URL.Query().Get("pid"))
	err := db.DeleteComment(id, pid)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		sendErrorResponse(w, def.Success)
	}
}

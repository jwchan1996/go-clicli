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
		sendMsg(w,401,"参数解析失败")
		return
	}

	if resp, err := db.AddComment(cbody.Content, cbody.Pid, cbody.Uid, cbody.Tuid, cbody.Vid, cbody.Time, cbody.Color); err != nil {
		log.Printf("%s", err)
		sendMsg(w,401,"数据库错误")
		return
	} else {
		res := def.Comment{Id: resp.Id, Content: resp.Content, Time: resp.Time, Pid: resp.Pid, Uid: resp.Uid}
		sendCommentResponse(w, res, 200)
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
		sendMsg(w,401,"数据库错误")
		log.Printf("%s", err)
		return
	} else {
		res := &def.Comments{Comments: resp}
		sendCommentsResponse(w, res, 200)
	}
}

func DeleteComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	AuthToken( w, r,3)
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	pid, _ := strconv.Atoi(r.URL.Query().Get("pid"))
	err := db.DeleteComment(id, pid)
	if err != nil {
		sendMsg(w,401,"数据库错误")
		return
	} else {
		sendMsg(w,200,"删除成功")
	}
}

package handler

import (
	"encoding/json"
	"github.com/132yse/acgzone-server/api/db"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
)

func AddVideo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	body := &def.Video{}

	if err := json.Unmarshal(req, body); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if resp, err := db.AddVideo(body.Oid, body.Title, body.Content, body.Pid, body.Uid); err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Video{Oid: resp.Oid, Title: resp.Title, Content: resp.Content, Time: resp.Time, Pid: resp.Pid, Uid: resp.Uid}
		sendVideoResponse(w, res, 201)
	}

}

func UpdateVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	vid, _ := strconv.Atoi(id)
	req, _ := ioutil.ReadAll(r.Body)
	body := &def.Video{}

	if err := json.Unmarshal(req, body); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if resp, err := db.UpdateVideo(vid, body.Oid, body.Title, body.Content, body.Pid, body.Uid); err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Video{Oid: resp.Oid, Title: resp.Title, Content: resp.Content, Time: resp.Time, Pid: resp.Pid, Uid: resp.Uid}
		sendVideoResponse(w, res, 201)
	}

}

func GetVideos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pid, _ := strconv.Atoi(r.URL.Query().Get("pid"))
	uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	resp, err := db.GetVideos(pid, uid, page, pageSize)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := &def.Videos{Videos: resp}
		sendVideosResponse(w, res, 201)
	}
}

func GetVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	vid, _ := strconv.Atoi(p.ByName("id"))
	resp, err := db.GetVideo(vid)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Video{Id: resp.Id, Oid: resp.Oid, Title: resp.Title, Content: resp.Content, Time: resp.Time, Pid: resp.Pid, Ptitle: resp.Ptitle, Uid: resp.Uid, Uname: resp.Uname, Uqq: resp.Uqq}
		sendVideoResponse(w, res, 201)
	}
}

func DeleteVideo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	pid, _ := strconv.Atoi(r.URL.Query().Get("pid"))

	err := db.DeleteVideo(id, pid)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		sendErrorResponse(w, def.Success)
	}
}

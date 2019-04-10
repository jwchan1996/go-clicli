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

func AddPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//role := RightAuth(w, r, p)
	//if role == "user" {
	//	sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
	//	return
	//}
	req, _ := ioutil.ReadAll(r.Body)
	pbody := &def.Post{}

	if err := json.Unmarshal(req, pbody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	resp, err := db.AddPost(pbody.Title, pbody.Content, pbody.Status, pbody.Sort, pbody.Tag, pbody.Uid)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Post{Id: resp.Id, Title: resp.Title, Content: resp.Content, Status: resp.Status, Sort: resp.Sort, Tag: resp.Tag, Time: resp.Time, Uid: resp.Uid}
		sendPostResponse(w, res, 201)
	}

}

func UpdatePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//role := RightAuth(w, r, p)
	//if role != "admin" || role != "editor" {
	//	sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
	//	return
	//}
	pid := p.ByName("id")
	pint, _ := strconv.Atoi(pid)
	req, _ := ioutil.ReadAll(r.Body)
	pbody := &def.Post{}
	if err := json.Unmarshal(req, pbody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if resp, err := db.UpdatePost(pint, pbody.Title, pbody.Content, pbody.Status, pbody.Sort, pbody.Tag, pbody.Time); err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Post{Id: resp.Id, Title: resp.Title, Content: resp.Content, Status: resp.Status, Sort: resp.Sort, Tag: resp.Tag, Time: resp.Time}
		sendPostResponse(w, res, 201)
	}

}

func DeletePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	role := RightAuth(w, r, p)
	if role != "admin" && role != "editor" {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}
	pid, _ := strconv.Atoi(p.ByName("id"))
	err := db.DeletePost(pid)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		sendErrorResponse(w, def.Success)
	}
}

func GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Cross(w)
	pid, _ := strconv.Atoi(p.ByName("id"))
	resp, err := db.GetPost(pid)
	if err != nil {
		log.Printf("%s",err)
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Post{Id: resp.Id, Title: resp.Title, Content: resp.Content, Status: resp.Status, Sort: resp.Sort, Tag: resp.Tag, Time: resp.Time, Uid: resp.Uid, Uname: resp.Uname, Uqq: resp.Uqq}
		sendPostResponse(w, res, 201)
	}
}

func GetPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Cross(w)
	status := r.URL.Query().Get("status")
	sort := r.URL.Query().Get("sort")
	tag := r.URL.Query().Get("tag")
	uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	resp, err := db.GetPosts(page, pageSize, status, sort, tag, uid)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := &def.Posts{Posts: resp}
		sendPostsResponse(w, res, 201)
	}
}

func SearchPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Cross(w)
	key := r.URL.Query().Get("key")

	resp, err := db.SearchPosts(key)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		log.Printf("%s", err)
		return
	} else {
		res := &def.Posts{Posts: resp}
		sendPostsResponse(w, res, 201)
	}

}

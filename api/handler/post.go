package handler

import (
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/132yse/acgzone-server/api/db"
	"log"
)

func AddPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	pbody := &def.Post{}

	if err := json.Unmarshal(req, pbody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if resp, err := db.AddPost(pbody.Title, pbody.Content, pbody.Status, pbody.Sort, pbody.Uid); err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Post{Id: resp.Id, Title: resp.Title, Content: resp.Content, Status: resp.Status, Sort: resp.Sort, Time: resp.Time, Uid: resp.Uid}
		sendPostResponse(w, res, 201)
	}

}

func UpdatePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := p.ByName("id")
	pint, _ := strconv.Atoi(pid)

	req, _ := ioutil.ReadAll(r.Body)
	pbody := &def.Post{}

	if err := json.Unmarshal(req, pbody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if resp, err := db.UpdatePost(pint, pbody.Title, pbody.Content, pbody.Status, pbody.Sort); err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Post{Id: resp.Id, Title: resp.Title, Content: resp.Content, Status: resp.Status, Sort: resp.Sort, Time: resp.Time}
		sendPostResponse(w, res, 201)
	}

}

func DeletePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	pid, _ := strconv.Atoi(p.ByName("id"))
	resp, err := db.GetPost(pid)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Post{Id: resp.Id, Title: resp.Title, Content: resp.Content, Status: resp.Status, Sort: resp.Sort, Time: resp.Time, Uid: resp.Uid, Uname: resp.Uname, Uqq: resp.Uqq}
		sendPostResponse(w, res, 201)
	}
}

func GetPostsType(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status := r.URL.Query().Get("status")
	sort := r.URL.Query().Get("sort")
	uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	resp, err := db.GetPostsByOneOf(status, sort, uid, page, pageSize)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		log.Printf("%s", err)
		return
	} else {
		res := &def.Posts{Posts: resp}
		sendPostsResponse(w, res, 201)
	}
}

func GetPostsAlso(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	status := r.URL.Query().Get("status")
	sort := r.URL.Query().Get("sort")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	resp, err := db.GetPostsByStatusAndSort(status, sort,page, pageSize)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		log.Printf("%s", err)
		return
	} else {
		res := &def.Posts{Posts: resp}
		sendPostsResponse(w, res, 201)
	}
}

func SearchPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

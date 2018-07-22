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

func AddPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	pbody := &def.Post{}

	if err := json.Unmarshal(req, pbody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if resp, err := db.AddPost(pbody.Title, pbody.Content, pbody.Status, pbody.Sort); err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Post{Id: resp.Id, Title: resp.Title, Content: resp.Content, Status: resp.Status, Sort: resp.Sort, Time: resp.Time}
		sendPostResponse(w, res, 201)
	}

}

func GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := p.ByName("id")
	pint, err := strconv.Atoi(pid)
	if err != nil {
		return
	}
	resp, err := db.GetPost(pint)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Post{Id: resp.Id, Title: resp.Title, Content: resp.Content, Status: resp.Status, Sort: resp.Sort, Time: resp.Time}
		sendPostResponse(w, res, 201)
	}
}

func GetPosts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	status := r.URL.Query().Get("status")
	sort := r.URL.Query().Get("sort")

	resp, err := db.GetPosts(status, sort)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		log.Printf("%s", err)
		return
	} else {
		res := &def.Posts{Posts: resp}
		sendPostsResponse(w, res, 201)
	}
}

package handler

import (
	"io"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/132yse/acgzone-server/api/db"
)

func Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &def.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if res, _ := db.GetUser(ubody.Name); res != nil {
		sendErrorResponse(w, def.ErrorUserNameRepeated)
		return
	}

	if err := db.CreateUser(ubody.Name, ubody.Pwd, ubody.Role, ubody.QQ, ubody.Desc); err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		sendErrorResponse(w, def.Success)
	}

}

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("name")
	resp, err := db.GetUser(uname)
	if err != nil {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	} else {
		res := def.UserCredential{Id: resp.Id, Name: resp.Name, Role: resp.Role, QQ: resp.QQ, Desc: resp.Desc}
		sendNormalResponse(w, res, 201)
	}
}

func AllPosts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "all ariticles!")
}

func GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "get a post")
}

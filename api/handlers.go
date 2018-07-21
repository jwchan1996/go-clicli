package main

import (
	"io"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/132yse/acgzone-server/api/dbOpt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	//res = bytes.TrimPrefix(res, []byte{239, 187, 191})
	ubody := &def.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbOpt.CreateUser(ubody.Name, ubody.Pwd, ubody.Role, ubody.QQ, ubody.Desc); err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		sendNormalResponse(w, string(res), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("name")
	io.WriteString(w, uname)
}

func AllPosts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "all ariticles!")
}

func GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "get a post")
}

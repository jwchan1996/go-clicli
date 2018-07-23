package handler

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/132yse/acgzone-server/api/db"
	"strconv"
)

func Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	ubody := &def.UserCredential{}

	if err := json.Unmarshal(req, ubody); err != nil {
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

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	ubody := &def.UserCredential{}

	if err := json.Unmarshal(req, ubody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	resp, err := db.GetUser(ubody.Name)
	if err != nil || len(resp.Pwd) == 0 || ubody.Pwd != resp.Pwd {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	} else {
		res := def.UserCredential{Id: resp.Id, Name: resp.Name, Role: resp.Role, QQ: resp.QQ, Desc: resp.Desc}
		sendUserResponse(w, res, 201)
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid := p.ByName("id")
	pint, _ := strconv.Atoi(pid)

	req, _ := ioutil.ReadAll(r.Body)
	ubody := &def.UserCredential{}

	if err := json.Unmarshal(req, ubody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if res, _ := db.GetUser(ubody.Name); res != nil {
		sendErrorResponse(w, def.ErrorUserNameRepeated)
		return
	}

	if resp, err := db.UpdateUser(pint, ubody.Name, ubody.Pwd, ubody.Role, ubody.QQ, ubody.Desc); err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.UserCredential{Id: resp.Id, Name: resp.Name, Role: resp.Role, QQ: resp.QQ, Desc: resp.Desc}
		sendUserResponse(w, res, 201)
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
		sendUserResponse(w, res, 201)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	role := r.URL.Query().Get("role")

	resp, err := db.GetUsers(role)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := &def.Users{Users: resp}
		sendUsersResponse(w, res, 201)
	}
}

package handler

import (
	"encoding/json"
	"github.com/132yse/acgzone-server/api/db"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/132yse/acgzone-server/api/util"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
	"log"
)

const DOMAIN = "clicli.us"

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	ubody := &def.User{}

	if err := json.Unmarshal(req, ubody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if res, _ := db.GetUser(ubody.Name, 0); res != nil {
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

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	ubody := &def.User{}

	if err := json.Unmarshal(req, ubody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	resp, err := db.GetUser(ubody.Name, 0)
	pwd := util.Cipher(ubody.Pwd)

	if err != nil || len(resp.Pwd) == 0 || pwd != resp.Pwd {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	} else {
		qq := http.Cookie{Name: "uqq", Value: resp.QQ, Path: "/", MaxAge: 86400, Domain: DOMAIN}
		uid := http.Cookie{Name: "uid", Value: strconv.Itoa(resp.Id), Path: "/", MaxAge: 86400, Domain: DOMAIN}
		http.SetCookie(w, &qq)
		http.SetCookie(w, &uid)
		res := &def.User{Id: resp.Id, Name: resp.Name, Role: resp.Role, QQ: resp.QQ, Desc: resp.Desc}
		sendUserResponse(w, res, 201, "登陆成功啦！")
	}

}

func Logout(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	cookieId := http.Cookie{Name: "uid", Path: "/", Domain: DOMAIN, MaxAge: -1}
	cookieQq := http.Cookie{Name: "uqq", Path: "/", Domain: DOMAIN, MaxAge: -1}
	http.SetCookie(w, &cookieId)
	http.SetCookie(w, &cookieQq)
	sendErrorResponse(w, def.Success)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pint, _ := strconv.Atoi(p.ByName("id"))

	req, _ := ioutil.ReadAll(r.Body)
	ubody := &def.User{}

	if err := json.Unmarshal(req, ubody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	res, _ := db.GetUser("", pint)
	if res.Name != ubody.Name || res.Id == pint {
		if resp, err := db.UpdateUser(pint, ubody.Name, ubody.Pwd, ubody.Role, ubody.QQ, ubody.Desc); err != nil {
			sendErrorResponse(w, def.ErrorDB)
			return
		} else {
			ret := &def.User{Id: resp.Id, Name: resp.Name, Role: resp.Role, QQ: resp.QQ, Desc: resp.Desc}
			sendUserResponse(w, ret, 201, "更新成功啦")
		}
	} else {
		sendErrorResponse(w, def.ErrorUserNameRepeated)
		return
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid, _ := strconv.Atoi(p.ByName("id"))
	err := db.DeleteUser(uid)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		sendErrorResponse(w, def.Success)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Cross(w, r)
	uname := r.URL.Query().Get("uname")
	uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	resp, err := db.GetUser(uname, uid)
	if err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	}
	res := &def.User{Id: resp.Id, Name: resp.Name, Role: resp.Role, QQ: resp.QQ, Desc: resp.Desc}
	sendUserResponse(w, res, 201, "")

}

func GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Cross(w, r)
	role := r.URL.Query().Get("role")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	resp, err := db.GetUsers(role, page, pageSize)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := &def.Users{Users: resp}
		sendUsersResponse(w, res, 201)
	}
}

func SearchUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Cross(w, r)
	key := r.URL.Query().Get("key")

	resp, err := db.SearchUsers(key)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := &def.Users{Users: resp}
		sendUsersResponse(w, res, 201)
	}
}

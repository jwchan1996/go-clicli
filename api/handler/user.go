package handler

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/132yse/acgzone-server/api/db"
	"strconv"
	"github.com/132yse/acgzone-server/api/util"
	"encoding/base64"
)

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	ubody := &def.UserCredential{}

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
	ubody := &def.UserCredential{}

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
		res := def.UserCredential{Id: resp.Id, Name: resp.Name, Role: resp.Role, QQ: resp.QQ, Desc: resp.Desc}
		uanme := base64.StdEncoding.EncodeToString([]byte(resp.Name))
		token := util.CreateToken(resp.Id)
		cookieId := http.Cookie{Name: "uname", Value: uanme, Path: "/", Domain: "chinko.cc"}
		cookieToken := http.Cookie{Name: "token", Value: token, Path: "/", Domain: "chinko.cc"}
		cookieQq := http.Cookie{Name: "uqq", Value: strconv.Itoa(resp.QQ), Path: "/", Domain: "chinko.cc"}
		http.SetCookie(w, &cookieId)
		http.SetCookie(w, &cookieQq)
		http.SetCookie(w, &cookieToken)
		sendUserResponse(w, res, 201, "登陆成功啦！")
	}

}

func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookieId := http.Cookie{Name: "uname", Path: "/", Domain: "chinko.cc", MaxAge: -1}
	cookieQq := http.Cookie{Name: "uqq", Path: "/", Domain: "chinko.cc", MaxAge: -1}
	http.SetCookie(w, &cookieId)
	http.SetCookie(w, &cookieQq)
	sendErrorResponse(w, def.Success)
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

	resp, err := db.UpdateUser(pint, ubody.Name, ubody.Pwd, ubody.Role, ubody.QQ, ubody.Desc)
	token := util.ResolveToken(r.Header.Get("Token"))
	uqq, err := r.Cookie("uqq")
	qq, _ := strconv.Atoi(uqq.Name)
	if i := UserIsLogin(qq, token); i != 1 {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	}
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.UserCredential{Id: resp.Id, Name: resp.Name, Role: resp.Role, QQ: resp.QQ, Desc: resp.Desc}
		sendUserResponse(w, res, 201, "更新成功啦")
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uid, _ := strconv.Atoi(p.ByName("id"))
	err := db.DeleteUser(uid)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	}
	token := util.ResolveToken(r.Header.Get("Token"))
	uqq, err := r.Cookie("uqq")
	qq, _ := strconv.Atoi(uqq.Name)
	if i := UserIsLogin(qq, token); i != 1 {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	}
	sendErrorResponse(w, def.Success)

}

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := r.URL.Query().Get("uname")
	uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	resp, err := db.GetUser(uname, uid)
	if err != nil {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	} else {
		res := def.UserCredential{Id: resp.Id, Name: resp.Name, Role: resp.Role, QQ: resp.QQ, Desc: resp.Desc}
		sendUserResponse(w, res, 201, "")
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

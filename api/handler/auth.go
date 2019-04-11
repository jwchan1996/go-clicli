package handler

import (
	"encoding/base64"
	"github.com/132yse/acgzone-server/api/db"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/util"
	"net/http"
)

//登陆校验，只负责校验登陆与否
func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Cross(w,r)
	uqq, err := r.Cookie("uqq")
	if err != nil || uqq == nil {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	} else {
		sendErrorResponse(w, def.Success)
	}
}

//鉴权校验，负责判断是否具有编辑和审核权限
func RightAuth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) string {
	uname, _ := r.Cookie("uname")
	name, _ := base64.StdEncoding.DecodeString(uname.Value)

	resp, err := db.GetUser(string(name), 0)
	if err != nil {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return ""
	}
	token, _ := r.Cookie("token")                      //从 cookie 里取 token
	newToken := util.CreateToken(resp.Name, resp.Role) //服务端生成新的 token

	if token.Value == newToken { //已经登陆
		if resp.Role == "admin" || resp.Role == "editor" {
			return resp.Role
		}
	} else {
		return ""
	}
	return ""
}

func Cross(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
	w.Header().Add("content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Credentials","true")
}

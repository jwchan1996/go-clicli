package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/132yse/acgzone-server/api/util"
	"github.com/132yse/acgzone-server/api/db"
	"encoding/base64"
)

func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	uname, err := r.Cookie("uname")
	name, _ := base64.StdEncoding.DecodeString(uname.Value)
	if err != nil || uname == nil {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	} else {
		resp, err := db.GetUser(string(name), 0)
		if err != nil {
			sendErrorResponse(w, def.ErrorNotAuthUser)
			return
		} else {
			res := def.UserCredential{Id: resp.Id, Name: resp.Name, Role: resp.Role, QQ: resp.QQ, Desc: resp.Desc}
			sendUserResponse(w, res, 201, "")
		}
	}

}

func UserIsLogin(uqq int, token string) int {
	t := util.CreateToken(uqq)
	res := util.ResolveToken(t)
	if token == res {
		return 1
	} else {
		return 0
	}
}

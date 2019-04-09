package handler

import (
	"encoding/base64"
	"github.com/132yse/acgzone-server/api/db"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/util"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	uname, _ := r.Cookie("uname")
	name, _ := base64.StdEncoding.DecodeString(uname.Value)

	resp, err := db.GetUser(string(name), 0)
	if err != nil {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	}
	token, _ := r.Cookie("token") //从 cookie 里取 token
	newToken := util.CreateToken(resp.Name, resp.Role) //服务端生成新的 token

	if token.Value == newToken {
		res := &def.User{Id: resp.Id, Name: resp.Name, Role: resp.Role, QQ: resp.QQ, Desc: resp.Desc}
		sendUserResponse(w, res, 201, "")
	} else {
		sendErrorResponse(w, def.ErrorNotAuthUser)
	}

}

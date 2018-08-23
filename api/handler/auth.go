package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/132yse/acgzone-server/api/util"
	"encoding/base64"
	"github.com/132yse/acgzone-server/api/db"
)

func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	uname, err := r.Cookie("uname")
	if err != nil || uname == nil {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	}
	name, _ := base64.StdEncoding.DecodeString(uname.Value)

	resp, err := db.GetUser(string(name), 0)
	if err != nil {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	}
	token, err := r.Cookie("token")
	if err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
	}

	i := UserIsLogin(uname.Value, token.Value)
	if i != 1 {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	}
	res := &def.UserCredential{Id: resp.Id, Name: resp.Name, Role: resp.Role, QQ: resp.QQ, Desc: resp.Desc}
	sendUserResponse(w, res, 201, "")

}

func UserIsLogin(name string, token string) int {
	res := util.CreateToken(name)
	if res == token {
		return 1
	} else {
		return 0
	}
}

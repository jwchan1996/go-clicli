package handler

import (
	"encoding/base64"
	"fmt"
	"github.com/132yse/acgzone-server/api/db"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/132yse/acgzone-server/api/util"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
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
func userAuth(_ http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	uname, err := r.Cookie("uname")
	log.Printf("%s", uname)
	if err != nil || uname == nil {
		return err
	}
	token, err := r.Cookie("token")
	log.Printf("%s", token)
	if err != nil {
		return err
	}

	i := UserIsLogin(uname.Value, token.Value)
	if i != 1 {
		return fmt.Errorf("%s", "用户token不匹配")

	}
	return nil
}

func UserIsLogin(name string, token string) int {
	res := util.CreateToken(name, "admin")
	log.Printf("%s", res)
	if res == token {
		return 1
	} else {
		return 0
	}
}

package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/def"
	"github.com/132yse/acgzone-server/api/util"
	"strconv"
)

func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	token := util.ResolveToken(r.Header.Get("Token"))
	uqq, err := r.Cookie("uqq")
	qq, _ := strconv.Atoi(uqq.Name)
	if i := UserIsLogin(qq, token); i != 1 {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	}
	if err != nil || uqq == nil {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	} else {
		sendErrorResponse(w, def.Success)
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

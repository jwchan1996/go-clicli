package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/def"
)

func Auth(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := r.Cookie("uname")
	if err != nil {
		return
	} else {
		sendErrorResponse(w, def.Success)
	}

}

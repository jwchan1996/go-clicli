package handler

import (
	auth "github.com/nilslice/jwt"
	"net/http"
	"time"
	"encoding/json"
	"io"
	"github.com/julienschmidt/httprouter"
)

func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	claims := map[string]interface{}{"exp": time.Now().Add(time.Hour * 24).Unix(), "level": 1}
	token, err := auth.New(claims)
	if err != nil {
		return
	}

	resStr, _ := json.Marshal(struct {
		Token string `json:"token"`
	}{Token: token})

	io.WriteString(w, string(resStr))
}

func AuthToken(w http.ResponseWriter, r *http.Request, level int) {
	token := r.Header.Get("token")
	if auth.Passes(token) {
		s := auth.GetClaims(token)
		if int(s["level"].(float64)) < level {
			io.WriteString(w, string("权限不足"))
		}
	} else {
		io.WriteString(w, string("token无效或过期"))
	}

}

package handler

import (
	auth "github.com/nilslice/jwt"
	"net/http"
	"time"
	"encoding/json"
	"io"
	"github.com/julienschmidt/httprouter"
)

//登陆校验，只负责校验登陆与否
func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	claims := map[string]interface{}{"exp": time.Now().Add(time.Hour * 24).Unix(), "level": 1}
	token, err := auth.New(claims)
	if err != nil {
		return
	}
	w.Header().Add("Authorization", "Bearer "+token)

	w.WriteHeader(http.StatusOK)
	resStr, _ := json.Marshal(struct {
		Token string `json:"token"`
	}{Token: token})

	io.WriteString(w, string(resStr))
}

func Cross(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Max-Age", "3600")
	w.Header().Add("Access-Control-Allow-Headers", "x-requested-with")
}

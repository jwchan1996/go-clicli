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
	//eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImV4cCI6MTU2Mzc4ODIwNiwiaWF0IjpudWxsLCJpc3MiOm51bGwsImp0aSI6bnVsbCwibGV2ZWwiOjEsIm5iZiI6bnVsbCwic3ViIjpudWxsfQ.aivx_tzi2SygsQM_RztVnxv76_470vxoScvD5nwtQoc
	//w.Header().Add("Authorization", "Bearer "+token)
	//w.WriteHeader(http.StatusOK)

	t := http.Cookie{Name: "token", Value: token, Path: "/", MaxAge: 86400, Domain: DOMAIN}
	http.SetCookie(w, &t)
	resStr, _ := json.Marshal(struct {
		Token string `json:"token"`
	}{Token: token})

	io.WriteString(w, string(resStr))
}
package handler

import (
	"io"
	"net/http"
	"encoding/json"
	"github.com/132yse/acgzone-server/api/def"
)

func sendErrorResponse(w http.ResponseWriter, errRes def.ErrorResponse) {
	resStr, _ := json.Marshal(&errRes)
	io.WriteString(w, string(resStr))
}

func sendUserResponse(w http.ResponseWriter, uRes def.UserCredential, sc int, msg string) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(struct {
		Code int                `json:"code"`
		Msg  string             `json:"msg,omitempty"`
		User def.UserCredential `json:"user"`
	}{sc, msg, uRes})

	io.WriteString(w, string(resStr))

}

func sendPostResponse(w http.ResponseWriter, pRes def.Post, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(struct {
		Code   int      `json:"code"`
		Result def.Post `json:"result"`
	}{sc, pRes})

	io.WriteString(w, string(resStr))
}

func sendPostsResponse(w http.ResponseWriter, pRes *def.Posts, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(struct {
		Code int `json:"code"`
		*def.Posts
	}{sc, pRes})

	io.WriteString(w, string(resStr))
}

func sendUsersResponse(w http.ResponseWriter, pRes *def.Users, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(struct {
		Code int `json:"code"`
		*def.Users
	}{sc, pRes})

	io.WriteString(w, string(resStr))
}

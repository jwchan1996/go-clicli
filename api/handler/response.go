package handler

import (
	"encoding/json"
	"github.com/cliclitv/go-clicli/api/def"
	"io"
	"net/http"
)

func sendUserResponse(w http.ResponseWriter, uRes *def.User, sc int, msg string) {
	w.WriteHeader(sc)

	resStr, _ := json.Marshal(struct {
		Code int      `json:"code"`
		Msg  string   `json:"msg,omitempty"`
		User def.User `json:"user"`
	}{sc, msg, *uRes})

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

func sendCommentResponse(w http.ResponseWriter, cRes def.Comment, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(struct {
		Code   int         `json:"code"`
		Result def.Comment `json:"result"`
	}{sc, cRes})

	io.WriteString(w, string(resStr))
}

func sendVideoResponse(w http.ResponseWriter, Res def.Video, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(struct {
		Code   int       `json:"code"`
		Result def.Video `json:"result"`
	}{sc, Res})

	io.WriteString(w, string(resStr))
}

func sendCommentsResponse(w http.ResponseWriter, pRes *def.Comments, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(struct {
		Code int `json:"code"`
		*def.Comments
	}{sc, pRes})

	io.WriteString(w, string(resStr))
}

func sendVideosResponse(w http.ResponseWriter, Res *def.Videos, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(struct {
		Code int `json:"code"`
		*def.Videos
	}{sc, Res})

	io.WriteString(w, string(resStr))
}

func sendCookieResponse(w http.ResponseWriter, cRes def.Cookie, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(struct {
		Code   int        `json:"code"`
		Result def.Cookie `json:"result"`
	}{sc, cRes})

	io.WriteString(w, string(resStr))
}

func sendMsg(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	resStr, _ := json.Marshal(struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{Code: code, Msg: msg})

	io.WriteString(w, string(resStr))
}

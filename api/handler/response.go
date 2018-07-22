package handler

import (
	"io"
	"net/http"
	"encoding/json"
	"github.com/132yse/acgzone-server/api/def"
)

func sendErrorResponse(w http.ResponseWriter, errRes def.ErrorResponse) {
	w.WriteHeader(errRes.Code)
	resStr, _ := json.Marshal(&errRes)
	io.WriteString(w, string(resStr))
}

func sendUserResponse(w http.ResponseWriter, uRes def.UserCredential, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(struct {
		Code   int                `json:"code"`
		Result def.UserCredential `json:"result"`
	}{sc, uRes})
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
		Code   int        `json:"code"`
		Result *def.Posts `json:"result"`
	}{sc, pRes})
	io.WriteString(w, string(resStr))
}

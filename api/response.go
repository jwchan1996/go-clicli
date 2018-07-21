package main

import (
	"io"
	"net/http"
	"github.com/132yse/acgzone-server/api/def"
	"encoding/json"
)

func sendErrorResponse(w http.ResponseWriter, errRes def.ErrorResponse) {
	w.WriteHeader(errRes.Code)
	resStr, _ := json.Marshal(&errRes.Error)
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, res def.Success, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(&def.Success{Code: sc, Result: res})
	io.WriteString(w, string(resStr))
}

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

func sendNormalResponse(w http.ResponseWriter, uRes def.UserCredential, sc int) {
	w.WriteHeader(sc)
	resStr, _ := json.Marshal(&uRes)
	io.WriteString(w, string(resStr))
}

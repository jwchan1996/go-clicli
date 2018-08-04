package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"github.com/132yse/acgzone-server/api/def"
	"encoding/json"
	"github.com/132yse/acgzone-server/api/db"
	"log"
)

func AddComment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	cbody := &def.Comment{}


	if err := json.Unmarshal(req, cbody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if resp, err := db.AddComment(cbody.Content, cbody.Pid, cbody.Uid); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Comment{Id: resp.Id, Content: resp.Content, Time: resp.Time, Pid: resp.Pid, Uid: resp.Uid}
		sendCommentResponse(w, res, 201)
	}

}

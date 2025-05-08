package resp

import (
	"encoding/json"
	"github.com/alg-grindel/saveSocial/pkg/errno"
	"log"
	"net/http"
)

type Data map[string]interface{}

type BaseResp struct {
	Code        int64
	Message     string
	Description Data
}

func WriteJson(rw http.ResponseWriter, e errno.Errno, description Data) {
	resp := &BaseResp{
		Code:        e.Code,
		Message:     e.Message,
		Description: description,
	}
	err := json.NewEncoder(rw).Encode(&resp)
	if err != nil {
		log.Panicln("[ERROR] pkg.resp: can not unmarshal resp")
	}
}

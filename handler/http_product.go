package handler
import (
	//"fmt"
	"time"
	"strings"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/gwtony/gapi/log"
	"github.com/gwtony/gapi/utils"
	"github.com/gwtony/gapi/api"
	"github.com/gwtony/gapi/errors"
)

type AddHandler struct {
	eh  *EtcdHandler
	cw  *CronWorker
	ch  *CronHandler
	log log.Log
}

type DeleteHandler struct {
	eh  *EtcdHandler
	ch  *CronHandler
	log log.Log
}

type UpdateHandler struct {
	eh  *EtcdHandler
	ch  *CronHandler
	log log.Log
}

type ReadHandler struct {
	eh  *EtcdHandler
	log log.Log
}

func (h *AddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	admin := false

	if r.Method != "POST" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.log)
		return
	}

	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Read from body failed"), errors.BadRequestError, h.log)
		return
	}
	r.Body.Close()

	data := &CronMessage{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Parse from body failed"), errors.BadRequestError, h.log)
		return
	}
	h.log.Info("Add record request: (%s) from client: %s", string(result), r.RemoteAddr)

	//TODO: check cron format
	im := &IdMessage{Jobid: data.Jobid}
	imv, _ := json.Marshal(im)

	api.ReturnResponse(r, w, string(imv), h.log)
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.ReturnResponse(r, w, "", h.log)
}


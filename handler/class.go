package handler
import (
	//"fmt"
	//"time"
	//"strings"
	//"strconv"
	"io/ioutil"
	"encoding/json"
	"net/http"
	//"net/url"
	"github.com/gwtony/gapi/log"
	//"github.com/gwtony/gapi/utils"
	"github.com/gwtony/gapi/api"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/db"
	"github.com/gstdio/rrb-backend/structs"
)

type ClassGetAllHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type ClassInsertHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type ClassUpdateHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

func (h *ClassGetAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//admin := false

	if r.Method != "GET" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	h.Log.Info("Get class request from client: %s", r.RemoteAddr)

	result, err := h.Mc.ClassGetAll()
	//err := json.Unmarshal(result, &data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Get class failed"), errors.BadGatewayError, h.Log)
		return
	}

	rv, _ := json.Marshal(result)

	api.ReturnResponse(r, w, string(rv), h.Log)
}

func (h *ClassInsertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//admin := false

	if r.Method != "POST" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Read from body failed"), errors.BadRequestError, h.Log)
		return
	}
	r.Body.Close()

	data := &structs.Class{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Parse from body failed"), errors.BadRequestError, h.Log)
		return
	}
	h.Log.Info("Insert class request from client: %s", r.RemoteAddr)

	err = h.Mc.ClassInsert(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Insert class failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

func (h *ClassUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//admin := false

	if r.Method != "POST" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Read from body failed"), errors.BadRequestError, h.Log)
		return
	}
	r.Body.Close()

	data := &structs.Class{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Parse from body failed"), errors.BadRequestError, h.Log)
		return
	}
	h.Log.Info("Insert class request from client: %s", r.RemoteAddr)

	err = h.Mc.ClassUpdate(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Update class failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

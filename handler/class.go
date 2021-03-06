package handler
import (
	"net/http"
	"encoding/json"
	"github.com/gwtony/gapi/log"
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

type ClassDeleteHandler struct {
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

	data := &structs.Class{}
	msg, err := parseBody(r, data)
	if err != nil {
		api.ReturnError(r, w, msg, err, h.Log)
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

	data := &structs.Class{}
	msg, err := parseBody(r, data)
	if err != nil {
		api.ReturnError(r, w, msg, err, h.Log)
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

func (h *ClassDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	data := &structs.Class{}
	msg, err := parseBody(r, data)
	if err != nil {
		api.ReturnError(r, w, msg, err, h.Log)
		return
	}

	h.Log.Info("Delete class request from client: %s", r.RemoteAddr)

	err = h.Mc.ClassDelete(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Delete class failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

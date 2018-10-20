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

type SaleGetAllHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type SaleInsertHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type SaleUpdateHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type SaleDeleteHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type SaleTestHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

func (h *SaleGetAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	h.Log.Info("Get sale request from client: %s", r.RemoteAddr)

	result, err := h.Mc.SaleGetAll()
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Get sale failed"), errors.BadGatewayError, h.Log)
		return
	}

	rv, _ := json.Marshal(result)

	api.ReturnResponse(r, w, string(rv), h.Log)
}

func (h *SaleInsertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	data := &structs.Sale{}
	msg, err := parseBody(r, data)
	if err != nil {
		api.ReturnError(r, w, msg, err, h.Log)
		return
	}

	h.Log.Info("Insert sale request from client: %s", r.RemoteAddr)

	err = h.Mc.SaleInsert(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Insert class failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

func (h *SaleUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	data := &structs.Sale{}
	msg, err := parseBody(r, data)
	if err != nil {
		api.ReturnError(r, w, msg, err, h.Log)
		return
	}

	h.Log.Info("Insert sale request from client: %s", r.RemoteAddr)

	err = h.Mc.SaleUpdate(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Update sale failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

func (h *SaleDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	data := &structs.Sale{}
	msg, err := parseBody(r, data)
	if err != nil {
		api.ReturnError(r, w, msg, err, h.Log)
		return
	}

	h.Log.Info("Delete sale request from client: %s", r.RemoteAddr)

	err = h.Mc.SaleDelete(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Delete sale failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

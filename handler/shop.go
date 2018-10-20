package handler
import (
	"encoding/json"
	"net/http"
	"github.com/gwtony/gapi/log"
	"github.com/gwtony/gapi/api"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/db"
	"github.com/gstdio/rrb-backend/structs"
)

type ShopGetAllHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type ShopInsertHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type ShopUpdateHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type ShopDeleteHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

func (h *ShopGetAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	h.Log.Info("Get shop request from client: %s", r.RemoteAddr)
	result, err := h.Mc.ShopGetAll()
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Get shop failed"), errors.BadGatewayError, h.Log)
		return
	}

	rv, _ := json.Marshal(result)

	api.ReturnResponse(r, w, string(rv), h.Log)
}

func (h *ShopInsertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	data := &structs.Shop{}
	msg, err := parseBody(r, data)
	if err != nil {
		api.ReturnError(r, w, msg, err, h.Log)
		return
	}

	h.Log.Info("Insert shop request from client: %s", r.RemoteAddr)

	err = h.Mc.ShopInsert(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Insert class failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

func (h *ShopUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	data := &structs.Shop{}
	msg, err := parseBody(r, data)
	if err != nil {
		api.ReturnError(r, w, msg, err, h.Log)
		return
	}

	h.Log.Info("Insert shop request from client: %s", r.RemoteAddr)

	err = h.Mc.ShopUpdate(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Update shop failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

func (h *ShopDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	data := &structs.Shop{}
	msg, err := parseBody(r, data)
	if err != nil {
		api.ReturnError(r, w, msg, err, h.Log)
		return
	}

	h.Log.Info("Delete shop request from client: %s", r.RemoteAddr)

	err = h.Mc.ShopDelete(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Delete shop failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

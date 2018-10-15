package handler
import (
	"strconv"
	//"io/ioutil"
	"encoding/json"
	"net/url"
	"net/http"
	"github.com/gwtony/gapi/log"
	"github.com/gwtony/gapi/api"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/db"
)

type ProductGetBySubClassIdHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type ProductGetByShopIdHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

func (h *ProductGetBySubClassIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	h.Log.Info("Get product request from client: %s", r.RemoteAddr)
	u, _ := url.Parse(r.URL.String())
    values, _ := url.ParseQuery(u.RawQuery)
	h.Log.Debug(values)
	id, ok := values["id"]
	if !ok {
		api.ReturnError(r, w, errors.Jerror("No id in args"), errors.BadRequestError, h.Log)
		return
	}
	if len(id) == 0 {
		api.ReturnError(r, w, errors.Jerror("No id value in args"), errors.BadRequestError, h.Log)
		return
	}
	idnum, err := strconv.Atoi(id[0])
    if err != nil {
		api.ReturnError(r, w, errors.Jerror("Invalid id value"), errors.BadRequestError, h.Log)
		return
    }

	h.Log.Debug("id num is ", idnum)
	result, err := h.Mc.GetProductBySubClassId(idnum)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Get product failed"), errors.BadGatewayError, h.Log)
		return
	}

	rv, _ := json.Marshal(result)

	api.ReturnResponse(r, w, string(rv), h.Log)
}

func (h *ProductGetByShopIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	h.Log.Info("Get product request from client: %s", r.RemoteAddr)
	u, _ := url.Parse(r.URL.String())
    values, _ := url.ParseQuery(u.RawQuery)
	h.Log.Debug(values)
	id, ok := values["id"]
	if !ok {
		api.ReturnError(r, w, errors.Jerror("No id in args"), errors.BadRequestError, h.Log)
		return
	}
	if len(id) == 0 {
		api.ReturnError(r, w, errors.Jerror("No id value in args"), errors.BadRequestError, h.Log)
		return
	}
	idnum, err := strconv.Atoi(id[0])
    if err != nil {
		api.ReturnError(r, w, errors.Jerror("Invalid id value"), errors.BadRequestError, h.Log)
		return
    }

	result, err := h.Mc.GetProductByShopId(idnum)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Get product failed"), errors.BadGatewayError, h.Log)
		return
	}

	rv, _ := json.Marshal(result)

	api.ReturnResponse(r, w, string(rv), h.Log)
}

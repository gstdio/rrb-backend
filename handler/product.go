package handler
import (
	"strconv"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"net/http"
	"github.com/gwtony/gapi/log"
	"github.com/gwtony/gapi/api"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/db"
	"github.com/gstdio/rrb-backend/structs"
)

type ProductGetBySubClassIdHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type ProductGetByShopIdHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}
type ProductInsertHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}
type ProductUpdateHandler struct {
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
	result, err := h.Mc.ProductGetBySubClassId(idnum)
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

	result, err := h.Mc.ProductGetByShopId(idnum)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Get product failed"), errors.BadGatewayError, h.Log)
		return
	}

	rv, _ := json.Marshal(result)

	api.ReturnResponse(r, w, string(rv), h.Log)
}

func (h *ProductInsertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	data := &structs.Product{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Parse from body failed"), errors.BadRequestError, h.Log)
		return
	}
	h.Log.Info("Insert product request from client: %s", r.RemoteAddr)

	err = h.Mc.ProductInsert(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Insert class failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

func (h *ProductUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	data := &structs.Product{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Parse from body failed"), errors.BadRequestError, h.Log)
		return
	}
	h.Log.Info("Insert product request from client: %s", r.RemoteAddr)

	err = h.Mc.ProductUpdate(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Update product failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

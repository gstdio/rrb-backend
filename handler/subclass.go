package handler
import (
	"strconv"
	"encoding/json"
	"net/url"
	"net/http"
	"github.com/gwtony/gapi/log"
	"github.com/gwtony/gapi/api"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/db"
	"github.com/gstdio/rrb-backend/structs"
)

type SubClassGetAllHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type SubClassGetByClassIdHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type SubClassInsertHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type SubClassUpdateHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

type SubClassDeleteHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

func (h *SubClassGetAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	h.Log.Info("Get subclass request from client: %s", r.RemoteAddr)

	result, err := h.Mc.SubClassGetAll()
	//err := json.Unmarshal(result, &data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Get subclass failed"), errors.BadGatewayError, h.Log)
		return
	}

	rv, _ := json.Marshal(result)

	api.ReturnResponse(r, w, string(rv), h.Log)
}

func (h *SubClassGetByClassIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	h.Log.Info("Get subclass request from client: %s", r.RemoteAddr)
	u, _ := url.Parse(r.URL.String())
    values, _ := url.ParseQuery(u.RawQuery)
	h.Log.Debug(values)
	id, ok := values["class_id"]
	if !ok {
		api.ReturnError(r, w, errors.Jerror("No class_id in args"), errors.BadRequestError, h.Log)
		return
	}
	if len(id) == 0 {
		api.ReturnError(r, w, errors.Jerror("No class_id value in args"), errors.BadRequestError, h.Log)
		return
	}
	idnum, err := strconv.Atoi(id[0])
    if err != nil {
		api.ReturnError(r, w, errors.Jerror("Id value invalid"), errors.BadRequestError, h.Log)
		return
    }

	result, err := h.Mc.SubClassGetByClassId(idnum)
	//err := json.Unmarshal(result, &data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Get subclass failed"), errors.BadGatewayError, h.Log)
		return
	}

	rv, _ := json.Marshal(result)

	api.ReturnResponse(r, w, string(rv), h.Log)
}

func (h *SubClassInsertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//admin := false

	if r.Method != "POST" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	data := &structs.SubClass{}
	msg, err := parseBody(r, data)
	if err != nil {
		api.ReturnError(r, w, msg, err, h.Log)
		return
	}

	h.Log.Info("Insert subclass request from client: %s", r.RemoteAddr)

	err = h.Mc.SubClassInsert(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Insert subclass failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

func (h *SubClassUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	data := &structs.SubClass{}
	msg, err := parseBody(r, data)
	if err != nil {
		api.ReturnError(r, w, msg, err, h.Log)
		return
	}

	h.Log.Info("Update class request from client: %s", r.RemoteAddr)

	err = h.Mc.SubClassUpdate(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Update subclass failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

func (h *SubClassDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	data := &structs.SubClass{}
	msg, err := parseBody(r, data)
	if err != nil {
		api.ReturnError(r, w, msg, err, h.Log)
		return
	}

	h.Log.Info("Delete class request from client: %s", r.RemoteAddr)

	err = h.Mc.SubClassDelete(data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Delete subclass failed"), errors.BadGatewayError, h.Log)
		return
	}

	api.ReturnResponse(r, w, "", h.Log)
}

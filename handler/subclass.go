package handler
import (
	//"fmt"
	//"time"
	//"strings"
	"strconv"
	//"io/ioutil"
	"encoding/json"
	"net/http"
	"net/url"
	"github.com/gwtony/gapi/log"
	"github.com/gwtony/gapi/api"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/db"
)

type SubClassGetAllHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}
type SubClassGetByClassIdHandler struct {
	Mc  *db.MysqlContext
	Log log.Log
}

func (h *SubClassGetAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		api.ReturnError(r, w, errors.Jerror("Method invalid"), errors.BadRequestError, h.Log)
		return
	}

	h.Log.Info("Get subclass request from client: %s", r.RemoteAddr)
	u, _ := url.Parse(r.URL.String())
    values, _ := url.ParseQuery(u.RawQuery)
	h.Log.Debug(values)

	result, err := h.Mc.GetAllSubClass()
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

	result, err := h.Mc.GetSubClassByClassId(idnum)
	//err := json.Unmarshal(result, &data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Get subclass failed"), errors.BadGatewayError, h.Log)
		return
	}

	rv, _ := json.Marshal(result)

	api.ReturnResponse(r, w, string(rv), h.Log)
}

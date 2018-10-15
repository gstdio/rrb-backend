package handler
import (
	//"fmt"
	//"time"
	//"strings"
	//"strconv"
	//"io/ioutil"
	"encoding/json"
	"net/http"
	//"net/url"
	"github.com/gwtony/gapi/log"
	//"github.com/gwtony/gapi/utils"
	"github.com/gwtony/gapi/api"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/db"
)

type ClassGetAllHandler struct {
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

	result, err := h.Mc.GetAllClass()
	//err := json.Unmarshal(result, &data)
	if err != nil {
		api.ReturnError(r, w, errors.Jerror("Get class failed"), errors.BadGatewayError, h.Log)
		return
	}

	rv, _ := json.Marshal(result)

	api.ReturnResponse(r, w, string(rv), h.Log)
}

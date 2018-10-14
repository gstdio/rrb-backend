package handler
import (
	//"fmt"
	//"time"
	//"strings"
	//"strconv"
	//"io/ioutil"
	//"encoding/json"
	"net/http"
	"github.com/gwtony/gapi/log"
	"github.com/gwtony/gapi/api"
	//"github.com/gwtony/gapi/errors"
)

type AddHandler struct {
	log log.Log
}

type DeleteHandler struct {
	log log.Log
}

func (h *AddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.ReturnResponse(r, w, "", h.log)
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.ReturnResponse(r, w, "", h.log)
}


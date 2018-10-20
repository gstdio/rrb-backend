package handler
import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/gwtony/gapi/errors"
)

func parseBody(r *http.Request, data interface{}) (string, error) {
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Jerror("Read from body failed"), errors.BadRequestError
	}
	r.Body.Close()

	err = json.Unmarshal(result, &data)
	if err != nil {
		return errors.Jerror("Parse from body failed"), errors.BadRequestError
	}
	return "", nil
}

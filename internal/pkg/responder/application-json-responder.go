package responder

import (
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cast"
)

type applicationJSONResponder struct{}

type jsonResponse struct {
	GUID string `json:"guid"`
}

func (responder applicationJSONResponder) ContentType() string {
	return "application/json"
}

func (responder applicationJSONResponder) WriteResponse(guid uuid.UUID, w *http.ResponseWriter) error {
	j := jsonResponse{GUID: guid.String()}
	raw, err := json.Marshal(j)
	if err != nil {
		return err
	}
	res, err := cast.ToStringE(raw)
	if err != nil {
		return err
	}

	fmt.Fprint(*w, res)
	return nil
}

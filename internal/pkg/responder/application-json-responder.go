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
	UUID string `json:"uuid"`
}

func (responder applicationJSONResponder) ContentType() string {
	return "application/json"
}

func (responder applicationJSONResponder) WriteResponse(uuid uuid.UUID, w *http.ResponseWriter) error {
	j := jsonResponse{UUID: uuid.String()}
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

package responder

import (
	"encoding/xml"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cast"
)

type applicationXMLResponder struct {
}

type xmlResponse struct {
	XMLName xml.Name `xml:"gimme-an-uuid"`
	UUID    string   `xml:"uuid"`
}

func (responder applicationXMLResponder) ContentType() string {
	return "application/xml"
}

func (responder applicationXMLResponder) WriteResponse(uuid uuid.UUID, w *http.ResponseWriter) error {
	j := xmlResponse{UUID: uuid.String()}
	raw, err := xml.Marshal(j)
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

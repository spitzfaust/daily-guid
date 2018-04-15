package responder

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type textPlainResponder struct{}

func (responder textPlainResponder) ContentType() string {
	return "text/plain"
}

func (responder textPlainResponder) WriteResponse(guid uuid.UUID, w *http.ResponseWriter) error {
	fmt.Fprint(*w, guid.String())
	return nil
}

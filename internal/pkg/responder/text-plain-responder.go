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

func (responder textPlainResponder) WriteResponse(uuid uuid.UUID, w *http.ResponseWriter) error {
	fmt.Fprint(*w, uuid.String())
	return nil
}

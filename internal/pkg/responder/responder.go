package responder

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// Responder is used to write different MIME types to a http response.
type Responder interface {
	WriteResponse(guid uuid.UUID, w *http.ResponseWriter) error
	ContentType() string
}

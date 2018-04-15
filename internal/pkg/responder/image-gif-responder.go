package responder

import (
	"image/gif"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/spitzfaust/daily-guid/internal/pkg/image"
)

type imageGIFResponder struct {
}

func (responder imageGIFResponder) ContentType() string {
	return "image/gif"
}

func (responder imageGIFResponder) WriteResponse(guid uuid.UUID, w *http.ResponseWriter) error {
	img := image.GenerateGUIDImage(guid)
	err := gif.Encode(*w, img, &gif.Options{NumColors: 256})
	if err != nil {
		return err
	}
	return nil
}

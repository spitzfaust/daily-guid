package responder

import (
	"image/png"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/spitzfaust/daily-guid/internal/pkg/image"
)

type imagePNGResponder struct {
}

func (responder imagePNGResponder) ContentType() string {
	return "image/png"
}

func (responder imagePNGResponder) WriteResponse(guid uuid.UUID, w *http.ResponseWriter) error {
	img := image.GenerateGUIDImage(guid)
	err := png.Encode(*w, img)
	if err != nil {
		return err
	}
	return nil
}

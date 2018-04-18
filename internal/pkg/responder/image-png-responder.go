package responder

import (
	"image/png"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/spitzfaust/gimme-an-uuid/internal/pkg/imagegen"
)

type imagePNGResponder struct {
}

func (responder imagePNGResponder) ContentType() string {
	return "image/png"
}

func (responder imagePNGResponder) WriteResponse(uuid uuid.UUID, w *http.ResponseWriter) error {
	img := imagegen.GenerateUUIDImage(uuid)
	err := png.Encode(*w, img)
	if err != nil {
		return err
	}
	return nil
}

package dailyguid

import (
	"net/http"
	"strings"

	"github.com/spitzfaust/daily-guid/internal/pkg/contexts"
	"github.com/spitzfaust/daily-guid/internal/pkg/responder"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spitzfaust/daily-guid/internal/pkg/data"
	"github.com/spitzfaust/daily-guid/internal/pkg/logging"
)

// mainHandler is the handler for requests to the application.
func mainHandler(w http.ResponseWriter, r *http.Request) {
	guid, err := data.GetDailyGUID()
	if err != nil {
		logging.LogWithContext(jww.ERROR, contexts.GetIDFromContext(r.Context()), "Could not get daily guid: %q", err)
		http.Error(w, "An error occurred 😿", http.StatusInternalServerError)
		return
	}
	acceptHeader := r.Header.Get("Accept")
	acceptHeaderTokens := strings.Split(acceptHeader, ",")
	var contentType string
	if len(acceptHeaderTokens) == 0 {
		contentType = "text/plain"
	} else {
		contentType = acceptHeaderTokens[0]
	}
	factory := responder.NewFactory()
	responder := factory.Create(contentType)
	if responder == nil {
		logging.LogWithContext(jww.ERROR, contexts.GetIDFromContext(r.Context()), "Requested content type %q is not supported.", contentType)
		http.Error(w, "Requested content type is not supported.", http.StatusNotAcceptable)
		return
	}
	w.Header().Set("Content-Type", responder.ContentType())
	err = responder.WriteResponse(*guid, &w)
	if err != nil {
		logging.LogWithContext(jww.ERROR, contexts.GetIDFromContext(r.Context()), "Could not write response for content type %q.", contentType)
		http.Error(w, "An error occurred 😿", http.StatusInternalServerError)
		return
	}
}

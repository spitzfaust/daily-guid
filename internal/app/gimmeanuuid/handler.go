package gimmeanuuid

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
	"github.com/spitzfaust/gimme-an-uuid/internal/pkg/contexts"
	"github.com/spitzfaust/gimme-an-uuid/internal/pkg/logging"
	"github.com/spitzfaust/gimme-an-uuid/internal/pkg/responder"

	"github.com/spf13/cast"
	jww "github.com/spf13/jwalterweatherman"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func extractNamespace(ps httprouter.Params) (*uuid.UUID, error) {
	namespaceParam := ps.ByName("namespace")
	namespace, err := uuid.FromString(namespaceParam)
	if err != nil {
		return nil, err
	}
	return &namespace, nil
}

func getV1UUID(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id := uuid.NewV1()
	generateUUIDResponse(w, r, id)
}
func getV2UUID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	domainParam := ps.ByName("domain")
	domain, err := cast.ToIntE(domainParam)
	if err != nil {
		logging.LogWithContext(jww.WARN, contexts.GetIDFromContext(r.Context()), "Domain %q is not an integer.", domainParam)
		http.Error(w, "Domain has to be an integer.", http.StatusBadRequest)
		return
	}
	id := uuid.NewV2(byte(domain))
	generateUUIDResponse(w, r, id)

}
func getV3UUID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	namespace, err := extractNamespace(ps)
	if err != nil {
		logging.LogWithContext(jww.WARN, contexts.GetIDFromContext(r.Context()), "Could not parse namespace: %q.", err)
		http.Error(w, "Could not parse namespace.", http.StatusBadRequest)
		return
	}
	name := ps.ByName("name")
	id := uuid.NewV3(*namespace, name)
	generateUUIDResponse(w, r, id)
}
func getV4UUID(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id := uuid.NewV4()
	generateUUIDResponse(w, r, id)
}
func getV5UUID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	namespace, err := extractNamespace(ps)
	if err != nil {
		logging.LogWithContext(jww.WARN, contexts.GetIDFromContext(r.Context()), "Could not parse namespace: %q.", err)
		http.Error(w, "Could not parse namespace.", http.StatusBadRequest)
		return
	}
	name := ps.ByName("name")
	id := uuid.NewV5(*namespace, name)
	generateUUIDResponse(w, r, id)
}

func generateUUIDResponse(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
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
	err := responder.WriteResponse(id, &w)
	if err != nil {
		logging.LogWithContext(jww.ERROR, contexts.GetIDFromContext(r.Context()), "Could not write response for content type %q.", contentType)
		http.Error(w, "An unexpected error occurred.", http.StatusInternalServerError)
		return
	}
}

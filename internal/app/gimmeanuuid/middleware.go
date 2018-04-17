package gimmeanuuid

import (
	"errors"
	"net/http"
	"time"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spitzfaust/gimme-an-uuid/internal/pkg/contexts"
	"github.com/spitzfaust/gimme-an-uuid/internal/pkg/logging"
)

// setRequestIDMiddleware sets the id value in the context of the request.
func setRequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := contexts.AddIDToContext(r.Context())
		requestWithContext := r.WithContext(ctx)
		next.ServeHTTP(w, requestWithContext)
	})
}

// logMiddleware logs the beginning and the end of a request.
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctxID := contexts.GetIDFromContext(r.Context())
		logging.LogWithContext(jww.INFO, ctxID, "Received request from %q for '%q'.", r.RemoteAddr, r.RequestURI)
		begin := time.Now()
		next.ServeHTTP(w, r)
		end := time.Now()
		processedIn := end.Sub(begin)
		logging.LogWithContext(jww.INFO, ctxID, "Processed request from %q for %q in %q.", r.RemoteAddr, r.RequestURI, processedIn)
	})
}

// disableCachingMiddleware sets headers to disable client side caching.
func disableCachingMiddleware(next http.Handler) http.Handler {
	var etagHeaders = []string{
		"ETag",
		"If-Modified-Since",
		"If-Match",
		"If-None-Match",
		"If-Range",
		"If-Unmodified-Since",
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, h := range etagHeaders {
			w.Header().Del(h)
		}
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		next.ServeHTTP(w, r)
	})
}

// recoverMiddleware logs fatal errors (panics) and returns an internal server error status code.
func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			rec := recover()
			if rec != nil {
				switch t := rec.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				jww.FATAL.Println(err)
				http.Error(w, "An error occurred 😿", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

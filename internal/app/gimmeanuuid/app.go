package gimmeanuuid

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	jww "github.com/spf13/jwalterweatherman"
)

// Run starts the server on the given port.
func Run(name, port string, logThreshold jww.Threshold) {
	jww.SetStdoutThreshold(logThreshold)
	jww.INFO.Printf("Starting %q", name)
	jww.INFO.Printf("Serving on port: %q", port)
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/api/uuid/v1", getV1UUID)
	router.GET("/api/uuid/v2/:domain", getV2UUID)
	router.GET("/api/uuid/v3/:namespace/:name", getV3UUID)
	router.GET("/api/uuid/v4", getV4UUID)
	router.GET("/api/uuid/v5/:namespace/:name", getV5UUID)
	middlewareChain := alice.New(recoverMiddleware, setRequestIDMiddleware, logMiddleware, disableCachingMiddleware).Then(router)
	err := http.ListenAndServe(":"+port, middlewareChain)
	jww.FATAL.Printf("%q encountered a fatal error: %q\n", name, err)
	os.Exit(1)
}

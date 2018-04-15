package dailyguid

import (
	"net/http"

	"github.com/justinas/alice"
	jww "github.com/spf13/jwalterweatherman"
)

// Run starts the server on the given port.
func Run(name, port string, logThreshold jww.Threshold) {
	jww.SetStdoutThreshold(logThreshold)
	jww.INFO.Printf("Starting %q", name)
	jww.INFO.Printf("Serving on port: %q", port)
	myHandler := http.HandlerFunc(mainHandler)
	chain := alice.New(recoverMiddleware, setRequestIDMiddleware, logMiddleware, disableCachingMiddleware).Then(myHandler)
	err := http.ListenAndServe(":"+port, chain)
	jww.FATAL.Printf("%q encountered a fatal error: %q\n", name, err)
}

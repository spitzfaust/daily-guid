package main

import (
	"flag"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spitzfaust/daily-guid/internal/app/dailyguid"
)

// ProgramName is the program name.
const ProgramName = "daily-guid"

func main() {
	port := flag.String("p", "9999", "port to serve on")
	dailyguid.Run(ProgramName, *port, jww.LevelInfo)
}

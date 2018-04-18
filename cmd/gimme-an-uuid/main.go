package main

import (
	"flag"

	jww "github.com/spf13/jwalterweatherman"
	"github.com/spitzfaust/gimme-an-uuid/internal/app/gimmeanuuid"
)

// ProgramName is the program name.
const ProgramName = "gimme-an-uuid"

func main() {
	port := flag.String("p", "9999", "port to serve on")
	flag.Parse()
	gimmeanuuid.Run(ProgramName, *port, jww.LevelInfo)
}

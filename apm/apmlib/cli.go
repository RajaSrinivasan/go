package apmlib

import (
	"flag"
)

var Verbose bool
var List bool

func init() {
	flag.BoolVar(&Verbose, "verbose", true, "verbose")
	flag.BoolVar(&List, "list", false, "list the contents of the apm file")
	flag.Parse()
}

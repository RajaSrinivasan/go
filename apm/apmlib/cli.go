package apmlib

import (
	"flag"
)

var Verbose bool
var Extract bool
var List bool

func init() {
	flag.BoolVar(&Verbose, "verbose", true, "verbose")
	flag.BoolVar(&Extract, "extract", false, "extract the contents of the apm file")
	flag.BoolVar(&List, "list", false, "list the contents of the apm file")
	flag.Parse()
}

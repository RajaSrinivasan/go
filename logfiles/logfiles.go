package main

import (
	"flag"
	"path"

	"./loglib"
)

var verbose bool
var rlmid string

func init() {
	flag.BoolVar(&verbose, "verbose", false, "verbosity")
	flag.StringVar(&rlmid, "rlm", "unknown", "RLM Id")
	flag.Parse()
	loglib.Verbose = verbose
}

func main() {

	ffn := flag.Arg(0)
	fft, _ := loglib.DateOfLogfile(ffn)
	for fn := 0; fn < flag.NArg(); fn++ {
		loglib.AnalyzeFile(flag.Arg(fn))
	}
	t := "RLM " + rlmid + " " + fft.Format(" 2006-01-02")
	loglib.PlotCPUTemp(path.Base(ffn)+".png", t)
}

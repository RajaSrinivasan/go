package logfiles

import (
	"testing"
)

func TestAnalyze(t *testing.T) {
	Verbose = true
	Analyze("20180701.zip")
	Analyze("20180822_1022.log")
	Analyze("badname.txt")
}

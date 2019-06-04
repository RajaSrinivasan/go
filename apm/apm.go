package main

import (
	"fmt"

	"./apmlib"
)
// Demo feature
func main() {

	fmt.Printf("Hello pass %s\n", apmlib.CryptPassphrase())
	cont := apmlib.Load(apmlib.Arg())
	if apmlib.List {
		cont.Show()
	}
}

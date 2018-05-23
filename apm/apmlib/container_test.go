package apmlib

import (
	"fmt"
	"testing"
)

func TestUUID(arg *testing.T) {
	var pkg *Container
	pkg = New()
	fmt.Printf("%v\n", pkg.id)
}

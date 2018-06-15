package apmlib

import (
	"fmt"
	"testing"
)

func TestPasswords(arg *testing.T) {
	fmt.Println(SignPassphrase())
	fmt.Println(CryptPassphrase())
}

package apmlib

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
)

func (pkg *Container) Encrypt() {
	var out bytes.Buffer
	pass := "pass:" + CryptPassphrase()
	infilename := filepath.Join(pkg.workarea, pkg.Basename(), ".zip")
	cmd := exec.Command("openssl", "aes-256-cbc", "-pass", pass, "-in", infilename, "-out", pkg.filename)
	cmd.Stdout = &out
	cmd.Run()
	fmt.Printf("%s", out.String())
}

func (pkg *Container) Decrypt() {
	var out bytes.Buffer
	pass := "pass:" + CryptPassphrase()
	outfilename := filepath.Join(pkg.workarea, pkg.Basename(), ".zip")
	cmd := exec.Command("openssl", "aes-256-cbc", "-d", "-pass", pass, "-in", pkg.filename, "-out", outfilename)
	cmd.Stdout = &out
	cmd.Run()
	fmt.Printf("%s", out.String())
}

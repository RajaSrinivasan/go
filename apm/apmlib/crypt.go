package apmlib

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func checkDigest(fn string, dig string) error {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	hmd5 := md5.New()
	_, err = io.Copy(hmd5, f)
	if err != nil {
		log.Fatal(err)
		return err
	}
	calcdig := hex.EncodeToString(hmd5.Sum(nil))
	if calcdig == dig {
		return nil
	}
	return fmt.Errorf("signatures do not match for %s", fn)
}

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

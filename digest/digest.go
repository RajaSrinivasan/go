package main

import (
	"crypto/md5"
	"crypto/sha256"
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
	"path/filepath"
)

var verbose bool
var fmd5 bool
var fsha256 bool
var fcrc32 string
var all bool
var toplevel string

func GenerateDigest(fn string) {
	fp, _ := filepath.Abs(fn)
	fmt.Printf("%40s ----------------------\n", fp)
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if fmd5 {
		hmd5 := md5.New()
		_, err := io.Copy(hmd5, f)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%10s : %x\n", "MD5", hmd5.Sum(nil))
	}

	f.Seek(0, 0)
	if fsha256 {
		hsha256 := sha256.New()
		_, err := io.Copy(hsha256, f)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%10s : %x\n", "SHA256", hsha256.Sum(nil))
	}
	f.Seek(0, 0)
	if fcrc32 != "" {
		var poltable *crc32.Table
		switch fcrc32 {
		case "IEEE":
			poltable = crc32.MakeTable(crc32.IEEE)
		case "Castagnoli":
			poltable = crc32.MakeTable(crc32.Castagnoli)
		case "Koopman":
			poltable = crc32.MakeTable(crc32.Koopman)
		default:
			fmt.Println("Unknown CRC32 polynomial %s. Will use IEEE\n", fcrc32)
			poltable = crc32.IEEETable
		}
		hcrc32 := crc32.New(poltable)
		_, err := io.Copy(hcrc32, f)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%10s : %x\n", "CRC32 "+fcrc32, hcrc32.Sum(nil))
	}
}

func DigestAll(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		GenerateDigest(path)
	}
	return nil
}

func main() {
	flag.BoolVar(&verbose, "verbose", true, "verbose")
	flag.BoolVar(&fmd5, "md5", false, "generate md5 signature")
	flag.BoolVar(&fsha256, "sha256", false, "generate sha256 signature")
	flag.BoolVar(&all, "all", false, "generate all signatures")
	flag.StringVar(&toplevel, "walk", "", "walk this directory")
	flag.StringVar(&fcrc32, "crc32", "", "crc32 IEEE|Castagnoli|Koopman")
	flag.Parse()
	if all {
		fmd5 = true
		fsha256 = true
		fcrc32 = "IEEE"
	}
	if verbose {
		if fmd5 {
			fmt.Println("MD5 signature")
		}
		if fsha256 {
			fmt.Println("SHA256 signature")
		}
		if fcrc32 != "" {
			fmt.Println("CRC32 %s", fcrc32)
		}
		if toplevel != "" {
			fmt.Println("Will walk the dir %s", toplevel)
		}
	}

	if toplevel == "" {
		for i := 0; i < flag.NArg(); i++ {
			GenerateDigest(flag.Arg(i))
		}
	} else {
		filepath.Walk(toplevel, DigestAll)
	}
}

package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var verbose bool
var packfile string

func List(fn string) {
	fmt.Printf("File: %s\n", fn)
	reader, err := zip.OpenReader(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	for _, file := range reader.File {
		fmt.Printf("%s Size %d Compressed %d CRC %d\n",
			file.Name,
			file.UncompressedSize,
			file.CompressedSize,
			file.CRC32)
	}
}

func PackFiles() {

	fmt.Printf("Creating %s\n", packfile)

	newfile, err := ioutil.TempFile("..", "apm")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created temp file %s\n", newfile.Name())
	newname := newfile.Name()
	defer os.Remove(newname)

	zwriter := zip.NewWriter(newfile)
	//defer zwriter.Close()

	for argno := 0; argno < flag.NArg(); argno++ {
		fmt.Printf("Pack the file %s\n", flag.Arg(argno))
		cfile := flag.Arg(argno)
		content, err := os.Open(cfile)
		if err != nil {
			fmt.Printf("Unable to open %s", cfile)
		} else {
			defer content.Close()
		}
		cfilestat, _ := content.Stat()

		hdr, _ := zip.FileInfoHeader(cfilestat)
		hdr.Method = zip.Deflate

		writer, _ := zwriter.CreateHeader(hdr)

		_, err = io.Copy(writer, content)
		if err != nil {
			fmt.Printf("Error packing %s", cfile)
		}
	}
	zwriter.Close()
	newfile.Close()
	newoutfile, _ := os.Open(newname)
	outfile, _ := os.Create(packfile)
	io.Copy(outfile, newoutfile)
	newfile.Close()
	outfile.Close()
}

func ListPackage() {
	var out bytes.Buffer
	cmd := exec.Command("unzip", "-l", packfile)
	cmd.Stdout = &out
	_ = cmd.Run()
	fmt.Printf("%s", out.String())
}

func EncryptPackage() {
	var out bytes.Buffer
	cmd := exec.Command("openssl", "aes-256-cbc", "-pass", "pass:srini", "-in", packfile, "-out", packfile+".apm")
	cmd.Stdout = &out
	cmd.Run()
	fmt.Printf("%s", out.String())
}

func DecryptPackage() {
	var out bytes.Buffer
	cmd := exec.Command("openssl", "aes-256-cbc", "-d", "-pass", "pass:srini", "-in", packfile+".apm", "-out", packfile+".out")
	cmd.Stdout = &out
	cmd.Run()
	fmt.Printf("%s", out.String())
}

func main() {
	flag.BoolVar(&verbose, "verbose", true, "verbose")
	flag.StringVar(&packfile, "pack", "", "pack the files into a zip file of this name")
	flag.Parse()

	fmt.Println(packfile)
	if packfile == "" {
		for argno := 0; argno < flag.NArg(); argno++ {
			List(flag.Arg(argno))
		}
	} else {
		PackFiles()
		ListPackage()
		EncryptPackage()
		DecryptPackage()
	}
}

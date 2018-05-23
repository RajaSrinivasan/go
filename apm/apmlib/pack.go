package apmlib

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func (pkg *Container) Unpack() {
	zipfilename := filepath.Join(pkg.workarea, pkg.Basename(), ".zip")
	rdr, err := zip.OpenReader(zipfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer rdr.Close()

	for _, f := range rdr.File {
		infile, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		savefilename := filepath.Join(pkg.workarea, f.Name)
		ofile, _ := os.Create(savefilename)
		io.Copy(ofile, infile)
		infile.Close()
		ofile.Close()
	}
}

func (pkg *Container) Pack() {
	//packed := filepath.Join(pkg.workarea, pkg.Basename(), ".zip")
}

package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

var verbose bool

type Default struct {
	XMLName  xml.Name `xml:"default"`
	Sync     string   `xml:"sync-j,attr"`
	Revision string   `xml:"revision,attr"`
	Upstream string   `xml:"master,attr"`
}

type Remote struct {
	XMLName xml.Name `xml:"remote"`
	Fetch   string   `xml:"fetch,attr"`
	Name    string   `xml:"name,attr"`
}

type Copyfile struct {
	XMLName xml.Name `xml:"copyfile"`
	Dest    string   `xml:"dest,attr"`
	Src     string   `xml:"src,attr"`
}

type Project struct {
	XMLName  xml.Name `xml:"project"`
	Remote   string   `xml:"remote,attr"`
	Name     string   `xml:"name,attr"`
	Revision string   `xml:"revision,attr"`
	Path     string   `xml:"path,attr"`
	Copyfile Copyfile
}

type Manifest struct {
	XMLName  xml.Name  `xml:"manifest"`
	Default  Default   `xml:"default"`
	Remote   Remote    `xml:"remote"`
	Projects []Project `xml:"project"`
}

func GetCommitId(p string, rev string) string {
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	temppath := path.Join(wd, p)
	if verbose {
		fmt.Printf("Changing working dir to %s\n", temppath)
	}
	os.Chdir(temppath)
	if verbose {

	}
	out, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		if verbose {
			fmt.Printf("Error %v getting commit id\n", err)
		}
		return rev
	}
	return string(out)
}

func process(mfname string, sfname string) {
	mf, err := os.Open(mfname)
	if err != nil {
		panic(err)
	}
	if verbose {
		fmt.Printf("Opened %s\n", mfname)
	}
	defer mf.Close()

	sf, err := os.Create(sfname)
	if err != nil {
		panic(err)
	}
	defer sf.Close()

	mfdata, err := ioutil.ReadAll(mf)
	if err != nil {
		panic(err)
	}

	var manifest Manifest
	xml.Unmarshal(mfdata, &manifest)

	fmt.Printf("Repository: %s\n", manifest.Remote.Fetch)
	fmt.Fprintf(sf, "[components]\n")
	fmt.Fprintf(sf, "repositoryhost=%s\n", manifest.Remote.Fetch)
	for i := 0; i < len(manifest.Projects); i++ {
		if verbose {
			fmt.Println("Project: " + manifest.Projects[i].Name)
		}
		if len(manifest.Projects[i].Revision) == 0 {
			if verbose {
				fmt.Println("Inherit Revision from default")
			}
			manifest.Projects[i].Revision = manifest.Default.Revision
		}
		if verbose {
			fmt.Printf("\tRevision: %s\n", manifest.Projects[i].Revision)
		}
		cmtid := GetCommitId(manifest.Projects[i].Path, manifest.Projects[i].Revision)
		if verbose {
			fmt.Printf("\tCommitId: = %s\n", cmtid)
		}
		fmt.Fprintf(sf, "%s=%s", manifest.Projects[i].Name, cmtid)
	}
}

func main() {
	flag.BoolVar(&verbose, "verbose", false, "verbose")
	flag.Parse()

	manifestfile := flag.Arg(0)
	sysconffile := flag.Arg(1)

	if verbose {
		fmt.Printf("Manifest file %s System Config file %s\n", manifestfile, sysconffile)
	}

	process(manifestfile, sysconffile)
}

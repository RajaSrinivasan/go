package apmlib

import (
	"errors"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/alyu/configparser"
	"github.com/satori/go.uuid"
)

const fileName = "Manifest"

func checkFile(nm string, siz string, sig string) error {
	fstat, err := os.Stat(nm)
	if err != nil {
		if os.IsNotExist(err) {
			panic("Missing file " + nm)
		} else {
			panic(err)
		}
		return err
	}

	expsize, _ := strconv.Atoi(siz)
	if int64(expsize) != fstat.Size() {
		panic("File Size does not match " + nm)
	}

	err = checkDigest(nm, sig)
	if err != nil {
		panic("HASH signatures dont match " + nm)
		return err
	}

	return nil
}

func (pkg *Container) reviewFiles() error {

	sec, err := pkg.configuration.Section("packfiles")
	sigsec, _ := pkg.configuration.Section("signatures")
	destsec, _ := pkg.configuration.Section("destination")

	sigfileName := ""
	destfileName := ""

	filenames := sec.OptionNames()

	for _, fn := range filenames {
		fnfields := strings.Split(fn, ",")
		err = checkFile(path.Join(pkg.workarea, fnfields[0]), fnfields[2], fnfields[3])
		if err != nil {
			return err
		}

		if sigsec != nil {
			sigfileName = sigsec.ValueOf(fn)
			sigfilefields := strings.Split(sigfileName, ",")
			err = checkFile(path.Join(pkg.workarea, sigfilefields[0]), sigfilefields[2], sigfilefields[3])
			if err != nil {
				return err
			}
		}

		if destsec != nil {
			destfileName = destsec.ValueOf(fnfields[0])
		} else {
			panic("No destination name for " + fnfields[0])
			return errors.New(fnfields[0])
		}
		pkg.AddContent(fnfields[0], destfileName)
	}
	return nil
}

func (pkg *Container) reviewScripts() error {
	sec, err := pkg.configuration.Section("scripts")
	if err != nil {
		return err
	}
	var preinstall string
	if sec.Exists("preinstall") {
		preinstall = sec.ValueOf("preinstall")
	} else {
		preinstall = ""
	}
	var postinstall string
	if sec.Exists("postinstall") {
		postinstall = sec.ValueOf("postinstall")
	} else {
		postinstall = ""
	}
	pkg.SetScripts(preinstall, postinstall)
	return nil
}

func (pkg *Container) reviewPackage() error {

	sec, err := pkg.configuration.Section("package")
	if err != nil {
		return err
	}

	pkg.pkgtype = Value(sec.ValueOf("type"))
	pkg.origin = sec.ValueOf("host")
	pkg.id, _ = uuid.FromString(sec.ValueOf("id"))

	pkg.created, _ = time.Parse(time.ANSIC, sec.ValueOf("created"))

	if sec.Exists("subtype") {
		pkgSubType := sec.ValueOf("subtype")
		if pkgSubType == "wificonfig" {
			pkg.pkgtype = WIFICONFIG
		}
	}
	return nil
}

func (pkg *Container) ReadManifest() error {
	configparser.Delimiter = "="
	var err error
	pkg.configuration, err = configparser.Read(path.Join(pkg.workarea, fileName))
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = pkg.reviewPackage()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = pkg.reviewScripts()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = pkg.reviewFiles()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

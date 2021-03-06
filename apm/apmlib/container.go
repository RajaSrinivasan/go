package apmlib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/alyu/configparser"
	"github.com/satori/go.uuid"
)

type PackageType int

const (
	CONTAINERVERSION = 1
)

const (
	PERSONALITY PackageType = iota + 1
	SERVICE
	WIFICONFIG
	UNKNOWN
)

var pkgTypeNames = [...]string{
	"personality",
	"service",
	"personality.wificonfig",
}

func (pt PackageType) Image() string {
	return pkgTypeNames[pt-1]
}

func Value(pt string) PackageType {
	for idx, val := range pkgTypeNames {
		if val == pt {
			return PackageType(idx)
		}
	}
	return UNKNOWN
}

type Content struct {
	name        string
	destination string
	signature   string
}

type Container struct {
	filename      string
	version       int
	id            uuid.UUID
	pkgtype       PackageType
	origin        string    // host where it was generated
	created       time.Time // when created
	workarea      string
	configuration *configparser.Configuration
	contents      []Content
	preinstall    string
	postinstall   string
}

func init() {
}

func finalizer(f *Container) {
	os.RemoveAll(f.workarea)
}

func New() *Container {
	newcont := new(Container)
	newcont.workarea, _ = ioutil.TempDir("", "apm")
	newcont.contents = make([]Content, 1)
	runtime.SetFinalizer(newcont, finalizer)
	return newcont
}

func (pkg *Container) Basename() string {
	bn := filepath.Base(pkg.filename)
	return strings.TrimSuffix(bn, filepath.Ext(bn))
}

func Load(fn string) *Container {
	cont := New()
	cont.filename, _ = filepath.Abs(fn)
	cont.Decrypt()
	cont.Unpack()
	return cont
}

func Create(fn string, pt PackageType) *Container {
	cont := New()
	cont.version = CONTAINERVERSION
	cont.pkgtype = pt
	cont.id, _ = uuid.NewV4()
	cont.filename, _ = filepath.Abs(fn)
	return cont
}

func (pkg *Container) SetScripts(pre string, post string) {
	pkg.preinstall = pre
	pkg.postinstall = post
}

func (pkg *Container) AddContent(src string, dest string) {
	var newcont Content
	newcont.name = src
	newcont.destination = dest
	newcont.signature = ""
	pkg.contents = append(pkg.contents, newcont)
}

func (pkg *Container) showHeader() {
	fmt.Printf("Filename: %s\n", pkg.filename)
	fmt.Printf("Version:  %d\n", pkg.version)
	uuidtxt, _ := pkg.id.MarshalText()
	fmt.Printf("Id:       %v\n", uuidtxt)
}

func (pkg *Container) Show() {
	pkg.showHeader()
}

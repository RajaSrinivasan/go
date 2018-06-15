package apmlib

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

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
)

var pkgTypeNames = [...]string{
	"personality",
	"service",
	"personality.wificonfig",
}

func (pt PackageType) Image() string {
	return pkgTypeNames[pt-1]
}

type Content struct {
	name        string
	destination string
	signature   string
}

type Container struct {
	filename string
	version  int
	id       uuid.UUID
	pkgtype  PackageType
	origin   string    // host where it was generated
	created  time.Time // when created
	workarea string
	contents []Content
}

func init() {
}

func finalizer(f *Container) {
	os.RemoveAll(f.workarea)
}

func New() *Container {
	newcont := new(Container)
	newcont.workarea, _ = ioutil.TempDir("", "apm")
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

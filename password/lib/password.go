package password

import (
	"fmt"
	"io"
	"time"
	"os"
	"math/rand"
	"crypto/md5"
	"encoding/hex"
	"github.com/alyu/configparser"
)

const saltLength = 24

type Storage struct {
	filename string
	file     *configparser.Configuration
}

func generateSalt() string {
	timens := int64(time.Now().Nanosecond())
	src := rand.NewSource(timens)
	randp := rand.New(src)
	newsalt := make([]byte,saltLength)
	randp.Read(newsalt)
	return hex.EncodeToString(newsalt)
}

func generateHash(salt string, txt string) string {
    hasher := md5.New()
		var b []byte
	  io.WriteString(hasher,salt)
		io.WriteString(hasher,txt)
		hash := hasher.Sum(b)
		return hex.EncodeToString(hash)
}

func init() {

}

func (storage *Storage) Set(app string, username string, password string) {
	fmt.Printf("Set app=%s username=%s password=%s\n", app, username, password)

	sec, err := storage.file.Section(app)
	if err != nil {
		sec = storage.file.NewSection(app)
	}
  usersalt := generateSalt()
	usernameenc := generateHash( usersalt, username)
	passwordenc := generateHash( usersalt, password)

	sec.Add("usersalt",usersalt)
	sec.Add("username",usernameenc)
	sec.Add("password",passwordenc)

	sec.Add(username, password)
	configparser.Save(storage.file, storage.filename)
}

func (storage *Storage) Verify(app string, username string, password string) bool {
	fmt.Printf("Verify app=%s username=%s password=%s\n", app, username, password)

	sec,err := storage.file.Section(app)
	if err != nil {
		 fmt.Printf("Undefined application %s\n",app)
		 return false
	}

	salt := sec.ValueOf("usersalt")
	usernameenc := sec.ValueOf("username")
  checkusernameenc := generateHash(salt,username)
	if usernameenc != checkusernameenc {
		fmt.Printf("username does not match\n")
		return false
	}

	passwordenc := sec.ValueOf("password")
	checkpasswordenc := generateHash(salt,password)
	if passwordenc != checkpasswordenc {
		fmt.Printf("Password does not match\n")
		return false
	}
	return true
}

func New(filename string) *Storage {
	fmt.Printf("Setting the Config file %s\n", filename)
  file,err := os.Open(filename)
	if err != nil {
		fmt.Printf("File does not exist %s. creating\n",filename)
		file,err := os.Create(filename)
		if err != nil {
			panic(err)
		} else {
			fmt.Fprintf(file,"# Created \n")
			file.Close()
		}
	}
	file.Close()

	stor := new(Storage)
	stor.filename = filename

	tempfile, err := configparser.Read(filename)
	if err != nil {
     panic(err)
	}
	stor.file = tempfile
	return stor
}

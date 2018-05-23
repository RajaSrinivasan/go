package apmlib

import (
	"encoding/hex"
)

var cryptPP []byte //{0x64, 0xe6, 0x74, 0xc7, 0xf5, 0xcf, 0xd5}
var signPP []byte  //{0xd5, 0xcf, 0xf5, 0xc7, 0x74, 0xe6, 0x64}

// Y64e674c7f5cfd5
// Zd5cff5c774e664

func init() {
	cryptPP = []byte{0x64, 0xe6, 0x74, 0xc7, 0xf5, 0xcf, 0xd5}
	signPP = []byte{0xd5, 0xcf, 0xf5, 0xc7, 0x74, 0xe6, 0x64}
}

func CryptPassphrase() string {
	return "Y" + hex.EncodeToString(cryptPP)
}

func SignPassphrase() string {
	return "Z" + hex.EncodeToString(signPP)
}

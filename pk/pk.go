package main

import (
  "fmt"
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509"
  //"encoding/asn1"
  "encoding/gob"
  "encoding/pem"
  "os"
  "math/big"
)

func main() {

  reader := rand.Reader
  bitSize := 2048

  randintmax := big.NewInt(100756052263040186)
  var randint *big.Int

  for i := 0; i<10; i++ {
     randint,_ = rand.Int(reader,randintmax)
     fmt.Printf("%v\n",randint)
  }

  randint,_ = rand.Prime(reader,1024)
  fmt.Printf("%v\n",randint)
  
key, err := rsa.GenerateKey(reader, bitSize)
checkError(err)

publicKey := key.PublicKey

saveGobKey("private.key", key)
savePEMKey("private.pem", key)

saveGobKey("public.key", publicKey)
savePublicPEMKey("public.pem", publicKey)
}

func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
	//asn1Bytes, err := asn1.Marshal(pubkey)
  asn1Bytes,err := x509.MarshalPKIXPublicKey(&pubkey)
	checkError(err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	checkError(err)
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

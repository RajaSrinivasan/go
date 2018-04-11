package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func DumpFile(fn string, blocklen int) {
	fp, err := filepath.Abs(fn)
	if err != nil {
		fmt.Printf("File: %s Error %s\n", fn, err.Error())
		return
	}
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		fmt.Printf("No such file %s\n", fn)
		return
	}
	fmt.Printf("File: %s\n", fp)
	offset := 0
	fd, err := os.Open(fp)
	defer fd.Close()
	buffer := make([]byte, blocklen)
	for {
		buffer = buffer[:cap(buffer)]
		n, _ := fd.Read(buffer)
		if n == 0 {
			break
		}
		fmt.Printf("%08x ", offset)
		buffer = buffer[:n]
		fmt.Printf("%s", hex.EncodeToString(buffer))
		for i := n; i < blocklen; i++ {
			fmt.Print("  ")
		}
		fmt.Print(" ")
		for _, b := range buffer {
			if strconv.IsPrint(rune(b)) {
				fmt.Printf("%s", string(b))
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
		offset = offset + n
	}
}

func main() {
	var blocklen = flag.Int("b", 16, "Blocklength")
	flag.Parse()
	fmt.Printf("Blocklength is %d\n", *blocklen)

	for i := 0; i < flag.NArg(); i++ {
		DumpFile(flag.Arg(i), *blocklen)
	}
}

package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"hash/crc32"
	"strconv"
)

var from int
var to int

func main() {
	flag.Parse()
	from, _ = strconv.Atoi(flag.Arg(0))
	to, _ = strconv.Atoi(flag.Arg(1))
	//table := crc32.MakeTable(crc32.IEEE)
	buf := make([]byte, 2)
	var i uint16
	for i = uint16(from); i < uint16(to)+1; i++ {
		binary.LittleEndian.PutUint16(buf, i)
		fmt.Printf("%d %d\n", i, crc32.ChecksumIEEE(buf))
	}
}

package main

import (
	"flag"
	"fmt"
	"strconv"
	"unsafe"

	"./lib"
)

var verbose bool
var uint_bits int

func EvaluateByte(f string, t string) {
	fromval, _ := strconv.ParseUint(f, 10, 8)
	toval, _ := strconv.ParseUint(t, 10, 8)
	if verbose {
		fmt.Printf("Evaluate crc from %d to %d\n", fromval, toval)
	}
	var evalint uint8
	var crc uint16
	for evalint = uint8(fromval); evalint <= uint8(toval); evalint++ {
		crc = crc16.UpdateBlock(crc, unsafe.Pointer(&evalint), unsafe.Sizeof(evalint))
		fmt.Printf("%4d : 0x%04x\n", evalint, crc)
	}
}
func main() {
	flag.BoolVar(&verbose, "verbose", true, "verbose")
	flag.IntVar(&uint_bits, "bits", 8, "No of bits. 8|16|32|64")
	flag.Parse()

	switch uint_bits {
	case 8:
		fmt.Println("Unsigned byte")
		EvaluateByte(flag.Arg(0), flag.Arg(1))
	case 16:
		fmt.Println("Unsigned word")
	case 32:
		fmt.Println("Unsigned long")
	case 64:
		fmt.Println("Unsigned long long")
	default:
		fmt.Println("Only 8, 16, 32 or 64 bit integers please")
	}
}

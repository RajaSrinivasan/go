package crc16

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestAlpha(arg *testing.T) {
	var alpha string = "abcdefghijklmnopqrstuvwxyz"
	var crcval uint16 = 0
	crcval = Update(crcval, []byte(alpha))
	fmt.Printf("CRC of %s is %d (0x%04x)\n", alpha, crcval, crcval)
}

func Extract(ptr unsafe.Pointer, size uintptr) []byte {
	out := make([]byte, size)
	for i := range out {
		out[i] = *((*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(i))))
	}
	return out
}

func TestTable(arg *testing.T) {
	var crcval uint16 = 0
	crcval = Update(crcval, Extract(unsafe.Pointer(&CRCtable), unsafe.Sizeof(CRCtable)))
	fmt.Printf("CRC of the table is %d (0x%x)\n", crcval, crcval)
}

func TestTableBlock(arg *testing.T) {
	var crcval uint16 = 0
	crcval = UpdateBlock(crcval, unsafe.Pointer(&CRCtable), unsafe.Sizeof(CRCtable))
	fmt.Printf("CRC of the table as a block is %d (0x%x)\n", crcval, crcval)
}

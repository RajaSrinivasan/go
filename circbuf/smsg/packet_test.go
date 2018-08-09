package smsg

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestShowCounters(t *testing.T) {
	t.Log("ShowCounters")
	ShowCounters()
}

func TestAdd(t *testing.T) {
	t.Log("Add small packet")

	for i := 0; i < 21; i++ {
		Add(byte(i))
	}
	Add(0x55)
	for i := 0; i < 21; i++ {
		Add(byte(i))
	}
	Add(0x55)
	ShowCounters()
	resetCounters()
	for i := 0; i < 21; i++ {
		Add(byte(i))
	}
	Add(0x55)
	for i := 0; i < 21; i++ {
		Add(byte(i))
	}
	Add(0x55)
	ShowCounters()
	resetCounters()
	for i := 0; i < 5; i++ {
		Add(0x55)
	}
	for i := 0; i < 21; i++ {
		Add(byte(i))
	}
	Add(0x55)
	ShowCounters()

}

func TestGet(t *testing.T) {
	fmt.Printf("Get msg")
	resetCounters()
	for i := 0; i < 5; i++ {
		Add(0x55)
	}
	for i := 0; i < 21; i++ {
		Add(byte(i))
	}
	Add(0xaa)
	Add(0xaa)
	Add(0xaa)
	Add(0x55)
	Add(0x55)
	ShowCounters()
	msg, err := Get()
	if err != nil {
		_ = fmt.Errorf("Error getting the packed message %s", err)
		return
	}
	fmt.Printf("Message Length %d %s\n", len(msg), hex.EncodeToString(msg))
}

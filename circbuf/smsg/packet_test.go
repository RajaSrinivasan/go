package smsg

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"encoding/binary"

	"../crc16"
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
	t.Log("Get msg")
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

func generatePacket(s string) {
	fmt.Printf("sending Packet %s\n", s)
	var pktcrc uint16
	msgbytes := []byte(s)
	Add(startOfMessage)
	for _, b := range msgbytes {
		Add(b)
		pktcrc = crc16.UpdateCRC(pktcrc, b)
	}
	crcbytes := make([]byte, 2)
	binary.BigEndian.PutUint16(crcbytes, pktcrc)
	Add(crcbytes[0])
	Add(crcbytes[1])
	Add(startOfMessage)
	ShowCounters()
}

func generateDataStream() {
	for m := 0; m < 5; m++ {
		msg := fmt.Sprintf("Test Message %d\n", m)
		generatePacket(msg)
		time.Sleep(time.Second)
	}
	generatePacket("quit")
	time.Sleep(5 * time.Second)
}

func TestChannels(t *testing.T) {
	t.Logf("Test the channel interface")
	resetCounters()
	ch, _ := SetupChannel(1)
	go generateDataStream()
	for msgno := 0; ; msgno++ {
		//time.Sleep(2 * time.Second)
		msg := <-ch
		if "quit" == string(msg[:]) {
			return
		}
		fmt.Printf("Received %s\n", string(msg[:]))
		ShowCounters()
	}
}

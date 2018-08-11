package smsg

import (
	"errors"
	"fmt"

	"../crc16"
)

const startOfMessage = 0x55
const escapeChar = 0xaa

const maxMessageSize = 256

type messageState int

const (
	waiting messageState = iota
	started
	escaping
	complete
)

var state messageState
var buffer []byte
var msglen int

var bytesReceived int
var bytesDropped int
var bytesPayload int
var packetsReceived int
var overheadBytes int
var packetsDropped int

var channel chan []byte

func init() {
	state = waiting
	buffer = make([]byte, maxMessageSize)
}

// ShowCounters prints out the statistics of message assembly
func ShowCounters() {
	fmt.Printf("Bytes Received            : %d\n", bytesReceived)
	fmt.Printf("      Payload             : %d\n", bytesPayload)
	fmt.Printf("      Dropped             : %d\n", bytesDropped)
	fmt.Printf("      Overhead            : %d\n", overheadBytes)
	fmt.Printf("Packets Received          : %d\n", packetsReceived)
	fmt.Printf("        Dropped           : %d\n", packetsDropped)
}

func resetCounters() {
	bytesReceived = 0
	bytesDropped = 0
	bytesPayload = 0
	packetsReceived = 0
	overheadBytes = 0
	packetsDropped = 0
}

// Add adds a byte to the message
func Add(b byte) error {
	bytesReceived++
	if msglen >= cap(buffer) {
		bytesDropped++
		return errors.New("buffer overflow")
	}
	switch state {
	case waiting:
		if b == startOfMessage {
			overheadBytes++
			state = started
		} else {
			bytesDropped++
		}
	case started:
		if b == escapeChar {
			overheadBytes++
			state = escaping
		} else {
			if b == startOfMessage {
				overheadBytes++
				if msglen > 0 {
					packetsReceived++
					state = complete
					var pktcrc uint16
					pktcrc = crc16.Update(pktcrc, buffer[:msglen-2])
					fmt.Printf("Packet crc is %04x\n", pktcrc)
					fmt.Printf("Received CRC is %02x%02x\n", buffer[msglen-2], buffer[msglen-1])
					if channel != nil {
						if len(channel) < cap(channel) {
							m, _ := Get()
							channel <- m
						}
					}
				} else {

				}
			} else {
				bytesPayload++
				buffer[msglen] = b
				msglen++
			}
		}
	case escaping:
		if b == escapeChar || b == startOfMessage {
			bytesPayload++
			buffer[msglen] = b
			msglen++
			state = started
		} else {
			bytesDropped++
			state = started
			return errors.New("invalid escaped char")
		}
	case complete:
		if b == startOfMessage {
			state = started
		} else {
			state = waiting
		}
		bytesDropped++
		if msglen > 0 {
			packetsDropped++
			msglen = 0
		}
		return errors.New("buffer not emptied")
	}

	return nil
}

// Get returns the message as a byte array and prepares the buffer
// to receive the next message
func Get() ([]byte, error) {
	switch state {
	case complete:
		state = waiting
		temp := make([]byte, msglen-2)
		copy(temp, buffer[:msglen-2])
		msglen = 0
		return temp, nil
	default:
		return make([]byte, 0), errors.New("message not available")
	}
}

// SetupChannel creates a new channel to which packets are delivered when complete. The argument c
// is the capacity of the channel. The packet channel is a singleton - with subsequent calls returning
// the same channel
func SetupChannel(c int) (chan []byte, error) {
	if channel == nil {
		channel = make(chan []byte, c)
		return channel, nil
	}
	return channel, errors.New("previously created channel")
}

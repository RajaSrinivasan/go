package smsg

import (
	"errors"
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

func init() {
	state = waiting
	buffer = make([]byte, maxMessageSize)
}

// Add adds a byte to the message
func Add(b byte) error {
	if msglen >= cap(buffer) {
		return errors.New("buffer overflow")
	}
	switch state {
	case waiting:
		if b == startOfMessage {
			state = started
		}
	case started:
		if b == escapeChar {
			state = escaping
		} else {
			if b == startOfMessage {
				if msglen > 0 {
					state = complete
				}
			} else {
				buffer[msglen] = b
				msglen++
			}
		}
	case escaping:
		if b == escapeChar || b == startOfMessage {
			buffer[msglen] = b
			msglen++
		} else {
			return errors.New("invalid escaped char")
		}
	case complete:
		return errors.New("buffer not emptied")
	}

	return nil
}

// Get returns the message as a byte array and prepares the buffer
// to receive the next message
func Get() ([]byte, error) {
	switch state {
	case complete:
		msglen = 0
		state = waiting
		temp := make([]byte, msglen)
		copy(temp, buffer)
		return temp, nil
	default:
		return make([]byte, 0), errors.New("message not available")
	}
}

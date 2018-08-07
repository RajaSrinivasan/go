package circbuf

import (
	"testing"
)

func TestMake(t *testing.T) {
	t.Log("Testing Make buffer")
	newbuf := Make(128)
	newbuf.Show()
}

func TestAdd(t *testing.T) {
	t.Log("Testing Adds to Circular buffer")
	newbuf := Make(16)
	newbuf.Show()
	for i := 0; i < 10; i++ {
		newbuf.Add(byte(i))
	}
	newbuf.Show()
}

package circbuf

import (
	"fmt"
)

type Buffer struct {
	Head      int
	Tail      int
	Count     int
	Container []byte
}

func init() {
	fmt.Printf("Circular Buffer\n")
}

func Make(cap int) *Buffer {
	var buf = new(Buffer)
	buf.Container = make([]byte, cap)
	buf.Head = 0
	buf.Tail = 0
	buf.Count = 0
	return buf
}

func (buf *Buffer) Show() {
	fmt.Printf("Buffer Capacity %d\n", cap(buf.Container))
	fmt.Printf("Head %d Tail %d Count %d\n", buf.Head, buf.Tail, buf.Count)
}

func (buf *Buffer) Add(b byte) {
	if buf.Count >= cap(buf.Container) {
		panic("Buffer Overflow")
	}
	buf.Container[buf.Head] = b
	if buf.Head >= cap(buf.Container) {
		buf.Head = 0
	} else {
		buf.Head++
	}
	buf.Count++
}

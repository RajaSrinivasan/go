package logfiles

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	ser1 := New("Ser1")
	ser1.show()
}

func TestAdd(t *testing.T) {
	ser2 := New("Ser2")
	for i := 0; i < 100; i++ {
		newsamp := Sample{time.Now(), float32(i)}
		ser2.Add(newsamp)
		time.Sleep(1e8)
	}
	ser2.show()
}

package loglib

import (
	"fmt"
	"testing"
	"time"
)

func TestAddCPUTemp(t *testing.T) {
	for i := 0; i < 10; i++ {
		AddCPUTemp(time.Now(), float32(i))
		fmt.Printf("%d %f\n", i, float32(i))
	}
	ShowCPUTemp()
}

package logfiles

import (
	"fmt"
	"testing"
)

func TestShowValidItems(t *testing.T) {
	ShowValidItems()
}

func TestValidItem(t *testing.T) {
	fmt.Printf("Valid Item CS = %v\n", ValidItem("CS"))
	fmt.Printf("Valid Item CPUTemp = %v\n", ValidItem("CPUTemp"))
}

func TestIndex(t *testing.T) {
	fmt.Printf("Index of CS = %d\n", Index("CS"))
	fmt.Printf("Index of CPUTemp = %d\n", Index("CPUTemp"))
}

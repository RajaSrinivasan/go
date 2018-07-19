package logfiles

import (
	"fmt"
	"time"
)

const MIN_SAMPLE_COUNT = 3600 * 8

type Sample struct {
	At    time.Time
	Value float32
}

type Series struct {
	Name    string
	Min     float64
	Max     float64
	Samples []Sample
}

func New(nm string) *Series {
	var temp = new(Series)
	temp.Name = nm
	temp.Min = 0
	temp.Max = 100
	//temp.Samples = make([]Sample, MIN_SAMPLE_COUNT)
	return temp
}

func (ser *Series) Add(s Sample) {
	ser.Samples = append(ser.Samples, s)
}

func (ser *Series) SetRange(min float64, max float64) {
	ser.Min = min
	ser.Max = max
}
func (ser *Series) show() {
	fmt.Printf("Series %s Length %d Capacity %d\n", ser.Name, len(ser.Samples), cap(ser.Samples))
	for idx, samp := range ser.Samples {
		fmt.Printf("%000d : %v %f\n", idx, samp.At, samp.Value)
	}
}

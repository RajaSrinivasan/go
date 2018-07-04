package loglib

import (
	"fmt"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const SECS_IN_HOUR = 3600
const SAMPLES_PER_DAY = SECS_IN_HOUR * 24

type CPUTempSample struct {
	attime      time.Time
	temperature float32
}

var CPUTempSamples []CPUTempSample

func init() {
	//CPUTempSamples = make([]CPUTempSample)
}

func AddCPUTemp(at time.Time, val float32) {
	newval := CPUTempSample{at, val}
	CPUTempSamples = append(CPUTempSamples, newval)
}

func ShowCPUTemp() {
	for _, val := range CPUTempSamples {
		fmt.Printf("%v : %f\n", val.attime, val.temperature)
	}
}

func PlotCPUTemp(fn string, title string) {
	p, err := plot.New()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	p.Title.Text = title + " CPU Temperatures "
	p.X.Label.Text = "Time (secs)"
	p.Y.Label.Text = "Temperature deg C"
	p.Y.Min = 75
	p.Y.Max = 105

	if Verbose {
		fmt.Printf("Plotting %d Temperature Samples\n", len(CPUTempSamples))
	}
	pts := make(plotter.XYs, len(CPUTempSamples))
	for i := range pts {
		pts[i].X = float64(CPUTempSamples[i].attime.Sub(CPUTempSamples[0].attime) / time.Second)
		pts[i].Y = float64(CPUTempSamples[i].temperature)
	}

	err = plotutil.AddLinePoints(p, "Temp", pts)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 4*vg.Inch, fn); err != nil {
		panic(err)
	}

}

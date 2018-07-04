package loglib

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var Verbose bool = false

const (
	AM = 1 + iota
	AS
	CS
	SS
	VS
	AB
	IP
	FI
	FO
	FD
	CPU_Usage
	Up_Time
	Pct_Mem_Used_sys
	Mem_Used_sys
	Mem_Free_sys
	Mem_Used_rlc
	Processes
)

var itemStatusLine = [...]string{
	"AM=",
	"AS=",
	"CS=",
	"SS=",
	"VS=",
	"AB=",
	"IP=",
	"FI=",
	"FO=",
	"FD=",
	"CPU Usage:",
	"Up Time:",
	"% Mem Used(sys):",
	"Mem Used(sys):",
	"Mem Free(sys):",
	"Mem Used(rlc):",
	"Processes:"}

const TimeStampLength = 22

var nullTime time.Time
var cpuTemp *regexp.Regexp
var blankTimeStamp string

func init() {
	nullTime = time.Now()
	cpuTemp = regexp.MustCompile("sensord.*temp1: (.*) C")

	blankTimeStamp = strings.Repeat(" ", TimeStampLength)
}

func ExtractTimeStamp(line string) (time.Time, error) {
	timstr := line[0:TimeStampLength]
	val, err := time.Parse(time.StampMicro, timstr)
	if err != nil {
		return nullTime, err
	}
	return val, nil
}

func ExtractCPUTemp(line string) (bool, time.Time, float32) {
	if cpuTemp.MatchString(line) {
		at, _ := ExtractTimeStamp(line)
		ts := cpuTemp.FindStringSubmatch(line)
		tv, _ := strconv.ParseFloat(ts[1], 32)
		return true, at, float32(tv)
	}
	return false, nullTime, 0.0
}

func ConnectionStatusLine(line string) (bool, time.Time) {
	if strings.Contains(line, "Connection Status") {
		tv, _ := ExtractTimeStamp(line)
		return true, tv
	}
	return false, nullTime
}

func ConnectionStatusDetailLine(line string) (bool, int, string) {
	if line[0:TimeStampLength] == blankTimeStamp {
		for idx, item := range itemStatusLine {
			pos := strings.Index(line, item)
			if pos > 0 {
				return true, idx, line[pos+len(item) : len(line)]
			}
		}
	}
	return false, 0, ""
}

func AnalyzeFile(fn string) {

	file, err := os.Open(fn)
	if err != nil {
		fmt.Printf("Error: File: %s %s\n", fn, err)
		return
	}
	defer file.Close()

	ts, err := DateOfLogfile(fn)
	basetime := time.Now()
	if err != nil {
		fmt.Printf("File name syntax does not conform to logile name syntax\n")
	} else {
		basetime = ts
	}
	var firstTimeStampTemp time.Time = nullTime
	var connectionStatus bool = false
	var connectionStatusTime time.Time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if Verbose {
			fmt.Printf("%s\n", line)
		}
		if connectionStatus {
			found, valtype, valstr := ConnectionStatusDetailLine(line)
			if found {
				if Verbose {
					fmt.Printf("Time %v : Type : %s Value %s\n", connectionStatusTime, itemStatusLine[valtype], valstr)
				}
			} else {
				connectionStatus = false
			}
		} else {
			found, attime, temp := ExtractCPUTemp(line)
			if found {
				acttime := OffsetYear(attime, basetime)
				AddCPUTemp(acttime, temp)
				if firstTimeStampTemp == nullTime {
					firstTimeStampTemp = acttime
				}
				if Verbose {
					fmt.Printf("Time %v : Type : CPU Temp Value %f\n", OffsetYear(attime, basetime), temp)
				}
			} else {
				found, attime := ConnectionStatusLine(line)
				if found {
					connectionStatus = true
					connectionStatusTime = attime
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: File: %s %s", fn, err)
	}

}

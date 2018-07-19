package logfiles

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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

var itemNames = [...]string{
	"AM",
	"AS",
	"CS",
	"SS",
	"VS",
	"AB",
	"IP",
	"FI",
	"FO",
	"FD",
	"CPU_Usage",
	"Up_Time",
	"Pct_Mem_Used_sys",
	"Mem_Used_sys",
	"Mem_Free_sys",
	"Mem_Used_rlc",
	"Processes"}

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

var cpuTemp *regexp.Regexp
var blankTimeStamp string
var gatheredStats []*Series
var cpuTempStats *Series

func init() {
	nullTime = time.Now()
	cpuTemp = regexp.MustCompile("sensord.*temp1: (.*) C")
	blankTimeStamp = strings.Repeat(" ", TimeStampLength)
	gatheredStats = make([]*Series, Processes)
}

func ValidItem(nm string) bool {
	for _, opt := range itemNames {
		if nm == opt {
			return true
		}
	}
	return false
}
func ShowValidItems() {
	for _, opt := range itemNames {
		fmt.Printf("%s\n", opt)
	}
}
func Index(nm string) int {
	for i, opt := range itemNames {
		if nm == opt {
			return i + 1
		}
	}
	return 0
}

func SetupStats(nm string) {
	if ValidItem(nm) {
		gatheredStats[Index(nm)] = New(nm)
	} else {
		if nm == "CPUTemp" {
			cpuTempStats = New("CPUTemp")
			cpuTempStats.SetRange(75, 100)
		}
	}
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

func AnalyzeFile(rdr io.Reader, base time.Time) {

	var connectionStatus bool = false
	var connectionStatusTime time.Time
	var tempSample Sample

	scanner := bufio.NewScanner(rdr)
	for scanner.Scan() {
		line := scanner.Text()
		if Verbose {
			fmt.Printf("%s\n", line)
		}
		if connectionStatus {
			found, valtype, valstr := ConnectionStatusDetailLine(line)
			if found {
				if gatheredStats[valtype] != nil {
					tempSample.At = connectionStatusTime
					val, _ := strconv.ParseFloat(valstr, 32)
					tempSample.Value = float32(val)
					gatheredStats[valtype].Add(tempSample)
				}
				if Verbose {
					fmt.Printf("Time %v : Type : %s Value %s\n", connectionStatusTime, itemStatusLine[valtype], valstr)
				}
			} else {
				connectionStatus = false
			}
		} else {
			found, attime, temp := ExtractCPUTemp(line)
			if found {
				attime := OffsetYear(attime, base)
				if cpuTempStats != nil {
					tempSample.At = attime
					tempSample.Value = temp
					cpuTempStats.Add(tempSample)
				}
				if Verbose {
					fmt.Printf("Time %v : Type : CPU Temp Value %f\n", attime, temp)
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
}

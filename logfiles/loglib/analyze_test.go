package loglib

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestAnalyzeExtractTimeStamp(t *testing.T) {
	ts, err := ExtractTimeStamp("Jun 20 23:28:27.339991 RL00122 file_transfer.py[5826]: 20180620_2228stats.log")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", ts.Format(time.StampMicro))
	ts, err = ExtractTimeStamp("                       RL00122 file_transfer.py[5826]: 20180620_2228stats.log")
	if err != nil {
		fmt.Printf("Got expected failure %s\n", err)
	}
	fmt.Printf("%s\n", ts.Format(time.StampMicro))
}

func TestExtractCPUTemp(t *testing.T) {
	f, ts, v := ExtractCPUTemp("Jun 20 23:28:27.339991 RL00122 file_transfer.py[5826]: 20180620_2228stats.log")
	if !f {
		fmt.Printf("Did not find temperature. Was not expecting.\n")
	}
	f, ts, v = ExtractCPUTemp("Jun 20 23:28:15.281759 RL00122 sensord[430]:   temp1: 72.9 C")
	if f {
		fmt.Printf("Found a temp value %f at %v. Expecting 72.9\n", v, ts)
	}
}

func TestConnectionStatusLine(t *testing.T) {
	s, at := ConnectionStatusLine("Jun 20 23:28:27.339991 RL00122 file_transfer.py[5826]: 20180620_2228stats.log")
	if s {
		fmt.Printf("Was not expecting a Connection Status=yes\n")
		t.Fail()
	}
	s, at = ConnectionStatusLine("Jun 20 23:41:39.178841 RL00122 rlc[694]: RLM RL00122 Connection Status")
	if s {
		fmt.Printf("Connection Status line at %v\n", at)
	} else {
		fmt.Printf("Was expecting Connection Status=yes\n")
	}
}

func TestConnectionStatusDetailLine(t *testing.T) {
	s, tid, v := ConnectionStatusDetailLine("                                         % Mem Used(sys):39.05912")
	fmt.Printf("%v %d %s\n", s, tid, v)
	s, tid, v = ConnectionStatusDetailLine("                                         Mem Used(sys):  410812416 ")
	fmt.Printf("%v %d %s\n", s, tid, v)
	s, tid, v = ConnectionStatusDetailLine("                                         Mem Free(sys):  640958464")
	fmt.Printf("%v %d %s\n", s, tid, v)
	s, tid, v = ConnectionStatusDetailLine("                                         Mem Used(rlc):  48664  ")
	fmt.Printf("%v %d %s\n", s, tid, v)
	s, tid, v = ConnectionStatusDetailLine("                                         Processes: 128    ")
	fmt.Printf("%v %d %s\n", s, tid, v)
}

func TestAnalyzeFile(t *testing.T) {
	AnalyzeFile("../20180621_0028.log")
}

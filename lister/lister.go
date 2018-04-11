package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func listfile(fn string) {
	fmt.Printf("---------------------------------------- %s\n", fn)
	f, e := os.Open(fn)
	if e != nil {
		log.Fatal(e)
		return
	}
	defer f.Close()

	linenum := 0
	reader := bufio.NewReader(f)
	for {
		l, e := reader.ReadString('\n')
		//l, _, e := reader.ReadLine()
		if e != nil {
			l = l + "\n"
		}
		if e == io.EOF {
			l = l + "\n"
		}
		linenum++
		fmt.Printf("%04d : %s", linenum, l)
		if e == io.EOF {
			break
		}
	}
	fmt.Printf("%04d lines\n", linenum)
}
func main() {

	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		listfile(flag.Arg(i))
	}
}

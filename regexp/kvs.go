package main

import (
	"flag"
	"fmt"
	"regexp"
)

func main() {
	var section_name = regexp.MustCompile(`\x5b.+\x5d`)
	var key_name = regexp.MustCompile("[a-zA-Z]+")
	flag.Parse()
	for n := 0; n < flag.NArg(); n++ {
		m := section_name.MatchString(flag.Arg(n))
		fmt.Println(flag.Arg(n))
		fmt.Printf("\tSection Name : %v\n", m)
		fmt.Printf("\tKey Name : %v\n", key_name.MatchString(flag.Arg(n)))
		fmt.Printf("\t%v\n", key_name.FindStringIndex(flag.Arg(n)))
	}
}

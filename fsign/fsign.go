package main

import (
	"flag"
	"log"
)

var verbose bool
var oper_sign bool
var oper_dir bool

func init() {

	flag.BoolVar(&verbose, "verbose", true, "be more verbose")
	flag.BoolVar(&oper_sign, "sign", false, "sign the files. generate .sig for each. Default = verify")
	flag.BoolVar(&oper_dir, "dir", true, "arguments are directories. All files")
	flag.Parse()

	if verbose {
		if oper_sign {
			log.Println("Will sign the arguments. Generate .sig files as output")
		} else {
			log.Println("Will verify the signatures of the argument files. against the .sig files")
		}
		if oper_dir {
			log.Println("Arguments are directories. Will operate on contents")
		} else {
			log.Println("Arguments are files")
		}
	}
}

func sign_arg(arg) {
	setup_sign()
}

func verify_arg(arg) {
	setup_verify()
}

func process(arg string) {
	if verbose {
		log.Println(arg)
	}
	if oper_sign {
		sign_arg(arg)
	} else {
		verify_arg(arg)
	}
}

func main() {
	for i := 0; i < flag.NArg(); i++ {
		process(flag.Arg(i))
	}
}

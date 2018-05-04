package main

import (
	"flag"
	"fmt"

	"./lib"
)

var verbose bool
var optset bool
var optver bool
var config string
var app string
var role string

func ShowArguments() {
	fmt.Printf("Password storage file= %s\n", config)
	fmt.Printf("Application (section) name= %s\n", app)
	fmt.Printf("Verify Option = %v\n", optver)
	fmt.Printf("Set option = %v\n", optset)
}

func main() {

	flag.BoolVar(&verbose, "verbose", true, "verbose")
	flag.BoolVar(&optset, "set", false, "set the password")
	flag.BoolVar(&optver, "verify", false, "verify the password")
	flag.StringVar(&app, "app", "app", "application name")
	flag.StringVar(&role, "role" , "user" , "role for the user")
	flag.StringVar(&config, "config", "password.conf", "configuration file")

	flag.Parse()

	if verbose {
		ShowArguments()
	}

	if optver && optset {
		fmt.Printf("Use verify or set option. Not both\n")
	}
	if !optver && !optset {
		fmt.Printf("Specify verify or set option.\n")
	}

	pconfig := password.New(config)
	if optset {
		pconfig.Set(app, role, flag.Arg(0), flag.Arg(1))
	}

	if optver {
		ver := pconfig.Verify(app, role, flag.Arg(0), flag.Arg(1))
		fmt.Printf("App %s user %s password %s verify %v\n", app, flag.Arg(0), flag.Arg(1), ver)
	}
}

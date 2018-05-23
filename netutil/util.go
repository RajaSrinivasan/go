package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	flag.Parse()
	fmt.Println("Net Utility")
	fmt.Println("LookupHost")
	fmt.Println(net.LookupHost(flag.Arg(0)))
	fmt.Println("LookupIP")
	fmt.Println(net.LookupIP(flag.Arg(0)))
	fmt.Println("LookupMX")
	fmt.Println(net.LookupMX(flag.Arg(0)))
	fmt.Println("LookupNS")
	fmt.Println(net.LookupNS(flag.Arg(0)))

	fmt.Println("Lookup port")
	port, _ := net.LookupPort("tcp", "rtmps")
	fmt.Printf("rtmps %v\n", port)
	port, _ = net.LookupPort("tcp", "ssh")
	fmt.Printf("ssh %v\n", port)
	port, _ = net.LookupPort("tcp", "https")
	fmt.Printf("https %v\n", port)
	port, _ = net.LookupPort("tcp", "srp")
	fmt.Printf("rsvp-e2e-ignore %v\n", port)
}

package main

import (
	"flag"
	"fmt"
	"strconv"

	"./numlib"
)

func test(n int) {
	fmt.Printf("------------------------------Testing %d\n", n)
	fmt.Printf("Digits of %d = %v\n", n, numlib.DigitsOf(n))
	fmt.Printf("Value of %v = %d\n", numlib.DigitsOf(n), numlib.DecimalValueOf(numlib.DigitsOf(n)))
	fmt.Printf("Divisors of %d = %v\n", n, numlib.DivisorsOf(n))
	fmt.Printf("Factors of %d = %v\n", n, numlib.FactorsOf(n))
	fmt.Printf("Prime ? %d = %v\n", n, numlib.Prime(n))
	fmt.Printf("Perfect ? %d = %v\n", n, numlib.Perfect(n))
	fmt.Printf("Harshad ? %d = %v\n", n, numlib.Harshad(n))
}

func main() {

	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		val, _ := strconv.Atoi(flag.Arg(i))
		test(val)
	}

}

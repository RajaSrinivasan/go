package main

import (
	"flag"
	"fmt"
	"strconv"

	"./numlib"
)

var fibonacci bool

func test(n int) {
	if n > 0 {
		fmt.Printf("------------------------------Testing %d\n", n)
		fmt.Printf("Digits of %d = %v\n", n, numlib.DigitsOf(n))
		fmt.Printf("Value of %v = %d\n", numlib.DigitsOf(n), numlib.DecimalValueOf(numlib.DigitsOf(n)))
		fmt.Printf("Divisors of %d = %v\n", n, numlib.DivisorsOf(n))
		fmt.Printf("Factors of %d = %v\n", n, numlib.FactorsOf(n))
		fmt.Printf("Prime ? %d = %v\n", n, numlib.Prime(n))
		fmt.Printf("Perfect ? %d = %v\n", n, numlib.Perfect(n))
		fmt.Printf("Harshad ? %d = %v\n", n, numlib.Harshad(n))
		fmt.Printf("Happy ? %d = %v\n", n, numlib.Happy(n))
		fmt.Printf("Happier ? %d = %v\n", n, numlib.Happier(n))
	}
}

func enumerate() {
	longest := 3
	var N1 int
	var N2 int

	for n1 := 3; n1 < 121; n1 += 2 {
		for n2 := n1 + 2; n2 < 121; n2 += 2 {
			if numlib.MutualPrime(n1, n2) {
				fmt.Printf("%d %d %d %d\n", n1, n2, n1*n2, n2-n1)
				if n1*n2 > longest {
					N1 = n1
					N2 = n2
					longest = N1 * N2
				}
			}
		}
		fmt.Printf("%d %d %d %d\n", N1, N2, longest, N2-N1)
	}
}

func exec_fibonacci_tests() {
	fl := numlib.Fibonacci()
	for idx, val := range fl {
		fmt.Printf("%4d : %d\n", idx, val)
		test(val)
	}
}

func main() {
	flag.BoolVar(&fibonacci, "fibonacci", false, "fibonacci tests")
	flag.Parse()
	if fibonacci {
		exec_fibonacci_tests()
	}
	for i := 0; i < flag.NArg(); i++ {
		val, _ := strconv.Atoi(flag.Arg(i))
		test(val)
	}

	for i := 0; i < flag.NArg()-1; i++ {
		num1, _ := strconv.Atoi(flag.Arg(i))
		num2, _ := strconv.Atoi(flag.Arg(i + 1))
		fmt.Printf("GCD of %d and %d is %d\n", num1, num2, numlib.Gcd(num1, num2))
	}

}

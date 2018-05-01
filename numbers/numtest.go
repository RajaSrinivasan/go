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

	for i := 0; i < flag.NArg()-1; i++ {
		num1, _ := strconv.Atoi(flag.Arg(i))
		num2, _ := strconv.Atoi(flag.Arg(i + 1))
		fmt.Printf("GCD of %d and %d is %d\n", num1, num2, numlib.Gcd(num1, num2))
	}

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

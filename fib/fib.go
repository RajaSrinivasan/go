package main

import "fmt"

type NumType uint64

var f0, f1, f2 NumType

func Gcd(num1 NumType, num2 NumType) NumType {
	if num1 > num2 {
		return Gcd(num2, num1)
	} else if num1 == num2 {
		return num1
	} else if num2%num1 == 0 {
		return num1
	}

	return Gcd(num1, num2%num1)
}

func main() {
	var fibArray [100]NumType
	var fibno = 0
	f0 = 0
	f1 = 1
	fibArray[fibno] = f0
	fibno++
	fmt.Printf("%d) %d\n", fibno, f0)
	fibArray[fibno] = f1
	fibno++
	fmt.Printf("%d) %d\n", fibno, f1)
	for {
		f2 = f0 + f1
		if f2 < f1 {
			break
		}
		f0 = f1
		f1 = f2
		fibArray[fibno] = f2
		fibno++
		fmt.Printf("%d) %d\n", fibno, f1)
	}

	for fn1 := 1; fn1 < fibno; fn1++ {
		for fn2 := fn1 + 1; fn2 < fibno; fn2++ {
			fgcd := Gcd(fibArray[fn1], fibArray[fn2])
			if fgcd > 1 {
				fmt.Printf("Gcd of %d and %d is %d\n", fibArray[fn1], fibArray[fn2], fgcd)
			}
		}
	}
}

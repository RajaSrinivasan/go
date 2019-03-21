package numlib

import (
	"fmt"
	"math"
)

func Gcd(num1 int, num2 int) int {
	if num1 > num2 {
		return Gcd(num2, num1)
	} else if num1 == num2 {
		return num1
	} else if num2%num1 == 0 {
		return num1
	}

	return Gcd(num1, num2%num1)
}

func MutualPrime(num1 int, num2 int) bool {
	if Gcd(num1, num2) == 1 {
		return true
	}
	return false
}

type NumPair struct {
	N1 int
	N2 int
}

type SumCube struct {
	Sum int
	P1  NumPair
	P2  NumPair
}

type TaxicabNumber struct {
	SC1 SumCube
	SC2 SumCube
}

func find(arr []SumCube, sum int) (bool, int) {
	for idx, sc := range arr {
		if sc.Sum == sum {
			return true, idx
		}
	}
	return false, 0
}

func TaxiCabNumbers(max int) []TaxicabNumber {
	var sumCube SumCube
	var tcNum TaxicabNumber
	sumCubes := make([]SumCube, 1024)
	taxCabs := make([]TaxicabNumber, 0)

	maxNo := int(math.Pow(float64(math.MaxInt64), 1.0/3.0))
	maxNo = 1000
	for n1 := 1; n1 < maxNo && len(taxCabs) < max; n1++ {
		n1cube := Cube(n1)
		for n2 := 1; n2 < n1 && len(taxCabs) < max; n2++ {
			n2cube := Cube(n2)
			sumCube.Sum = n1cube + n2cube
			sumCube.P1 = NumPair{n1, n1cube}
			sumCube.P2 = NumPair{n2, n2cube}
			status, idx := find(sumCubes, sumCube.Sum)
			if status {
				tcNum.SC1 = sumCubes[idx]
				tcNum.SC2 = sumCube
				taxCabs = append(taxCabs, tcNum)
				//fmt.Printf("%5d) %10d ", len(taxCabs), sumCube.Sum)
				//PrintTaxicabNumber(tcNum)
			} else {
				sumCubes = append(sumCubes, sumCube)
			}
		}
	}
	return taxCabs
}
func showSumCube(sc SumCube) {
	fmt.Printf(" : %10d , %10d | %10d , %10d || ", sc.P1.N1, sc.P1.N2, sc.P2.N1, sc.P2.N2)
}

func PrintTaxicabNumber(tcNum TaxicabNumber) {
	//fmt.Printf("%10d = ", tcNum.SC1.Sum)
	showSumCube(tcNum.SC1)
	showSumCube(tcNum.SC2)
	fmt.Println()
}
func PrintTaxicabNumbers(tc []TaxicabNumber) {
	for idx, tcNum := range tc {
		fmt.Printf("%4d) ", idx)
		fmt.Printf("%10d || ", tcNum.SC1.Sum)
		PrintTaxicabNumber(tcNum)
	}
}

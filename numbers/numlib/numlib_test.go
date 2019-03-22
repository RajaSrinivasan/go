package numlib

import (
	"fmt"
	"testing"
)

func TestCubes(t *testing.T) {
	t.Log("Test Generate list of cubes")
	cubes := Powers(3)
	PrintArr("Generate Cubes", cubes)
}

func TestCube(t *testing.T) {
	t.Logf("Testing Cube\n")
	for i := 5; i < 15; i++ {
		ic := Cube(i)
		fmt.Printf("%d -> %d\n", i, ic)
	}
}

func TestSquare(t *testing.T) {
	t.Logf("Testing Squares\n")
	for i := 5; i < 15; i++ {
		ic := Square(i)
		fmt.Printf("%d -> %d\n", i, ic)
	}
}

func TestPrintSumCube(t *testing.T) {
	t.Logf("TestPrintSumCube\n")
	var sc SumCube
	sc.P1 = NumPair{5, Cube(5)}
	sc.P2 = NumPair{7, Cube(7)}
	sc.Sum = sc.P1.N2 + sc.P2.N2
	showSumCube(sc)
}

func TestPowers(t *testing.T) {

	for i := 10; i < 20; i++ {
		fmt.Printf("%4d ", i)
		for p := 2; p < 9; p++ {
			fmt.Printf(" %15d ", Power(i, p))
		}
		fmt.Println()
	}
}

func TestPrintTaxicabNumbers(t *testing.T) {
	var sc SumCube
	sc.P1 = NumPair{5, Cube(5)}
	sc.P2 = NumPair{7, Cube(7)}
	sc.Sum = sc.P1.N2 + sc.P2.N2
	var sc2 SumCube
	sc2.P1 = NumPair{6, Cube(6)}
	sc2.P2 = NumPair{9, Cube(9)}
	sc2.Sum = sc2.P1.N2 + sc2.P2.N2
	var tc TaxicabNumber
	tc.SC1 = sc
	tc.SC2 = sc2

	tcs := make([]TaxicabNumber, 0, 64)
	tcs = append(tcs, tc)
	PrintTaxicabNumbers(tcs)
}

func TestTaxicabNumbers(t *testing.T) {
	tcnums := TaxiCabNumbers(128)
	PrintTaxicabNumbers(tcnums)
}

func TestTaxicabNumbersOrder(t *testing.T) {
	tcnums := TaxiCabNumbersOrder(128, 4)
	PrintTaxicabNumbers(tcnums)
}

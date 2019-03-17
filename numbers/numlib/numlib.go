package numlib

import (
	"container/list"
	"fmt"
	"math"
	"sort"
)

func Odd(n int) bool {
	if n%2 == 0 {
		return false
	}
	return true
}
func Print(nm string, diglist *list.List) {
	fmt.Println(nm)
	for dig := diglist.Front(); dig != nil; dig = dig.Next() {
		fmt.Printf("%d ", dig.Value)
	}
	fmt.Println("")
}

func Convert(L *list.List) []int {
	result := make([]int, L.Len())
	elem := L.Front()
	for i := 0; i < L.Len(); i++ {
		result[i] = elem.Value.(int)
		elem = elem.Next()
	}
	return result
}

func SumOf(input []int) int {
	result := 0
	for _, val := range input {
		result += val
	}
	return result
}

func SumSqOf(input []int) int {
	result := 0
	for _, val := range input {
		result += val * val
	}
	return result
}

func ProductOf(input []int) int {
	result := 0
	for _, val := range input {
		result *= val
	}
	return result
}

func DecimalValueOf(input []int) int {
	result := 0
	for _, val := range input {
		result = result*10 + val
	}
	return result
}
func DigitsOf(n int) []int {
	nl := list.New()
	tempn := n
	for {
		d := tempn % 10
		nl.PushFront(d)
		tempn = tempn / 10
		if tempn == 0 {
			break
		}
	}
	return Convert(nl)
}

func DivisorsOf(n int) []int {
	nsqrt := int(math.Sqrt(float64(n))) + 1
	nl := list.New()
	for d := 1; d < nsqrt; d++ {
		if n%d == 0 {
			nl.PushBack(d)
			dd := n / d
			nl.PushBack(dd)
		}
	}
	result := Convert(nl)
	sort.Ints(result)
	return result
}

func FactorsOf(n int) []int {
	nsqrt := int(math.Sqrt(float64(n))) + 1
	nl := list.New()
	nl.PushFront(1)
	tempn := n
	for d := 2; d < nsqrt; {
		if tempn%d == 0 {
			nl.PushBack(d)
			tempn = tempn / d
			if tempn == 1 {
				break
			}
		} else {
			d++
		}
	}
	if tempn != 1 {
		nl.PushBack(tempn)
	}
	result := Convert(nl)
	sort.Ints(result)
	return result
}

func Fibonacci() []int {
	nl := list.New()
	f0 := 0
	nl.PushBack(f0)
	f1 := 1
	nl.PushBack(f1)
	for {
		f2 := f1 + f0
		if f2 < f1 {
			break
		}
		nl.PushBack(f2)
		f0 = f1
		f1 = f2
	}
	result := Convert(nl)
	return result
}

func Prime(n int) bool {
	f := DivisorsOf(n)
	if len(f) == 2 {
		return true
	}
	return false
}

func Perfect(n int) bool {
	ds := DivisorsOf(n)
	sumds := SumOf(ds)
	if sumds/2 == n {
		return true
	}
	return false
}

func Harshad(n int) bool {
	ds := DigitsOf(n)
	sumds := SumOf(ds)
	if n%sumds == 0 {
		return true
	}
	return false
}

func Happy(n int) bool {
	var seen [1000]bool

	numnow := n
	for attempt := 0; attempt < 1000; attempt = attempt + 1 {
		ds := DigitsOf(numnow)
		numnow = SumSqOf(ds)
		fmt.Printf("%04d : %d\n", attempt, numnow)
		if numnow == 1 {
			return true
		}
		if seen[numnow] {
			return false
		}
		seen[numnow] = true
	}
	return false
}

func Happier(n int) bool {
	var seen [1000]bool

	numnow := n
	for attempt := 0; attempt < 1000; attempt = attempt + 1 {
		ds := DigitsOf(numnow)
		numnow = SumSqOf(ds)
		//fmt.Printf("%04d : %d\n", attempt, numnow)
		if numnow == n {
			return true
		}
		if seen[numnow] {
			return false
		}
		seen[numnow] = true
	}
	return false
}

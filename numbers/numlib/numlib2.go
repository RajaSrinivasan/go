package numlib

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

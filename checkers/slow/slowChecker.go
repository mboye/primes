package slow

import "github.com/mboye/primes"

type slowPrimeChecker struct {
}

// New returns a new slow prime checker
func New() primes.Checker {
	return slowPrimeChecker{}
}

func (c slowPrimeChecker) IsPrime(value int) bool {
	if value == 2 {
		return true
	} else if value < 2 {
		return false
	}

	for divisor := 2; divisor < value; divisor++ {
		if value%divisor == 0 {
			return false
		}
	}

	return true
}

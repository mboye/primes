package fast

import (
	"github.com/mboye/primes"
	"github.com/mboye/primes/checkers/slow"
)

type fastPrimeChecker struct {
	knownPrimes []int
	slowChecker primes.Checker
}

// New returns a new fast prime checker
func New() primes.Checker {
	return &fastPrimeChecker{
		knownPrimes: []int{2},
		slowChecker: slow.New()}
}

func (c *fastPrimeChecker) IsPrime(value int) bool {
	if value < 2 {
		return false
	}

	for _, prime := range c.knownPrimes {
		if value == prime {
			return true
		}

		if value%prime == 0 {
			return false
		}
	}

	if c.slowChecker.IsPrime(value) {
		c.knownPrimes = append(c.knownPrimes, value)
		return true
	}
	return false
}

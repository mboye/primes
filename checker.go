package primes

// Checker checks if an input value is a prime number.
type Checker interface {
	IsPrime(value int) bool
}

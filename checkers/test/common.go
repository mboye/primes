package test

import (
	"testing"

	"github.com/mboye/primes"
	"github.com/stretchr/testify/assert"
)

// CommonIsPrimeTest tests an implementation of Checker
func CommonIsPrimeTest(t *testing.T, checker primes.Checker) {
	cases := []struct {
		value   int
		isPrime bool
	}{
		{-1, false},
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{197, true}}

	for _, tc := range cases {
		assert.Equal(t, tc.isPrime, checker.IsPrime(tc.value), "Input value: %d, expected output: %t", tc.value, tc.isPrime)
	}
}

// CommonIsPrimeBenchmark benchmarks an implementation of Checker
func CommonIsPrimeBenchmark(b *testing.B, checker primes.Checker) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			checker.IsPrime(j)
		}
	}
}

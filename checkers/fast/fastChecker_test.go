package fast

import (
	"testing"

	ct "github.com/mboye/primes/checkers/test"
)

func TestIsPrime(t *testing.T) {
	ct.CommonIsPrimeTest(t, New())
}

func BenchmarkIsPrime(b *testing.B) {
	ct.CommonIsPrimeBenchmark(b, New())
}

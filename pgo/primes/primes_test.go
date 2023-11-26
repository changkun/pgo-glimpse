package primes_test

import (
	"fmt"
	"testing"
)

func primeSieve(max int) []int {
	var primes []int
	isComposite := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		if !isComposite[i] {
			primes = append(primes, i)
			for j := i * i; j <= max; j += i {
				isComposite[j] = true
			}
		}
	}
	return primes
}

func BenchmarkPrimeSieve(b *testing.B) {
	for j := 5; j <= 10; j *= 2 {
		b.Run(fmt.Sprintf("%v", j), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = primeSieve(j)
			}
		})
	}
}

package main_test

import "testing"

func f(x int) {
	if x == 0 {
		return
	}

	f(x - 1)
}

func BenchmarkCallOverhead(b *testing.B) {
	b.Run("call", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f(1024)
		}
	})
	b.Run("direct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for i := 0; i < 1024; i++ {
				_ = 1
			}
		}
	})
}

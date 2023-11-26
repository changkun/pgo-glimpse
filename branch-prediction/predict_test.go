package branchprediction_test

import (
	"testing"

	_ "unsafe"
)

func predictableBranching(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			sum += i
		} else {
			sum += 2 * i
		}
	}
}

//go:linkname fastrand runtime.fastrand
func fastrand() uint32

func unpredictableBranching(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		if fastrand()%2 == 0 {
			sum += i
		} else {
			sum += 2 * i
		}
	}
}

func BenchmarkBranchPrediction(b *testing.B) {
	b.Run("predictable", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			predictableBranching(100_000_000)
		}
	})
	b.Run("unpredictable", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			unpredictableBranching(100_000_000)
		}
	})
}

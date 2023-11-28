package math_test

import (
	"fmt"
	"testing"

	"changkun.de/x/pgo-glimpse/pgo/math"
)

func BenchmarkMul(b *testing.B) {
	size := 1 << 10
	m1 := math.NewRandMat[float32](size, size)
	m2 := math.NewRandMat[float32](size, size)
	var outCPU math.Mat[float32]
	b.Run(fmt.Sprintf("CPU-block(%vx%v)", size, size), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			outCPU = m1.Mul(m2)
		}
	})
	_ = outCPU
}

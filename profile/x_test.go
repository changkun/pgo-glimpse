package profile_test

import (
	"testing"

	"changkun.de/x/pgo-glimpse/profile"
)

func BenchmarkCounter(b *testing.B) {
	c := profile.Counter{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Inc()
		}
	})
}

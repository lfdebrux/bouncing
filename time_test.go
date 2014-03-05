package bouncing

import "testing"

func BenchmarkVondrakZenith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VondrakZenith(&J{P:&P{Beta:2}})
	}
}

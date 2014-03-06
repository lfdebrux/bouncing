package bouncing

import "testing"

import "math"

func BenchmarkVondrakZenith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VondrakZenith(&J{P:&P{Beta:2}})
	}
}

func TestVondrakZenithPositiveTime(t *testing.T) {
	for beta := 0.0; beta < 2*math.Pi; beta += 0.1 {
		j := &J{P:&P{Beta:beta,Phi:math.Pi/2}}
		VondrakZenith(j)
		if j.Time < 0 {
			t.Fatalf("VondrakZenith is giving negative time to sunrise at beta=%v",beta)
		}
	}
}

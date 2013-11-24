package bouncing

import "testing"

import "math"

func TestRandParticleInRange(t *testing.T) {
	const NUM = 100000
	for i := 0; i < NUM; i++ {
		p := RandParticle(Water)
		if p.Phi >= math.Pi || 0 > p.Phi {
			t.Fatalf("particle latitude should be in range 0..Pi, instead got %f",p.Phi)
		}
		if p.Beta >= 2*math.Pi || 0 > p.Beta {
			t.Fatalf("particle longitude should be in range 0..2Pi, instead got %f",p.Beta)
		}
	}
}

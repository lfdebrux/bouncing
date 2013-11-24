package particle

import "testing"

import "math"

const NUM = 100000

func TestRandParticleInRange(t *testing.T) {
	for i := 0; i < NUM; i++ {
		p := RandParticle(Water)
		if p.Phi >= math.Pi || 0 > p.Phi {
			t.Fatalf("Particle latitude should be in range 0..Pi, instead got %f",p.Phi)
		}
		if p.Beta >= 2*math.Pi || 0 > p.Beta {
			t.Fatalf("Particle longitude should be in range 0..2Pi, instead got %f",p.Beta)
		}
	}
}

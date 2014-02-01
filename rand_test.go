package bouncing

import "testing"

import "math"

func makeVelJ() *J {
	return &J{P:new(P)}
}

func TestUnInitedMaxwellian(t *testing.T) {
	if Q != nil {
		FreeMaxwellian()
	}
	defer FreeMaxwellian()

	RandVelocity(makeVelJ())
}

func BenchmarkRandVelocity(b *testing.B) {
	InitMaxwellian()
	defer FreeMaxwellian()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RandVelocity(makeVelJ())
	}
}

func TestRandDirectionInRange(t *testing.T) {
	for i := 0; i < NUM; i++ {
		psi,thetadash := RandDirection() 
		if psi >= 2*math.Pi || 0 > psi {
			t.Errorf("psi should be in range 0..2pi, instead got %f",psi)
		}
		if thetadash >= math.Pi/2 || 0 > thetadash {
			t.Errorf("thetadash should be in range 0..pi/2, instead got %f",thetadash)
		}
	}
}

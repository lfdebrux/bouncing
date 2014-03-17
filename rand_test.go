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
		j := makeVelJ()
		RandDirection(j) 
		if j.Psi >= 2*math.Pi || 0 > j.Psi {
			t.Errorf("psi should be in range 0..2pi, instead got %f",j.Psi)
		}
		if j.ThetaDash >= math.Pi/2 || 0 > j.ThetaDash {
			t.Errorf("thetadash should be in range 0..pi/2, instead got %f",j.ThetaDash)
		}
	}
}

func TestRandInitialPositionButlerInRange(t *testing.T) {
	const num = 100000
	for i := 0; i < num; i++ {
		p := new(P)
		RandInitialPositionButler(p)
		if p.Phi >= math.Pi || 0 > p.Phi {
			t.Fatalf("RandInitialPositionButler(p) => p.Phi = %f Pi; want 0..Pi",p.Phi/math.Pi)
		}
		if p.Beta >= 2*math.Pi || 0 > p.Beta {
			t.Fatalf("RandInitialPositionButler(p) => p.Beta = %f Pi; want 0..2Pi",p.Beta/math.Pi)
		}
	}
}

func TestRandInitialPositionVondrakInRange(t *testing.T) {
	const num = 100000
	for i := 0; i < num; i++ {
		p := new(P)
		RandInitialPositionVondrak(p)
		if p.Phi >= math.Pi || 0 > p.Phi {
			t.Fatalf("RandInitialPositionVondrak(p) => p.Phi = %f Pi; want 0..Pi",p.Phi/math.Pi)
		}
		if math.Pi/2 < p.Beta && p.Beta < 3*math.Pi/2 {
			t.Fatalf("RandInitialPositionVondrak(p) => p.Beta = %f Pi; want 3Pi/2..Pi/2",p.Beta/math.Pi)
		}
	}
}
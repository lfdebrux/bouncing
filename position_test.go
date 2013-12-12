package bouncing

import "testing"

import "math"

func BenchmarkPositionJump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ButlerPositionJump(1.35,3.1,463,2.98,1.7)
	}
}

func BenchmarkVondrakPositionJump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VondrakPositionJump(1.35,3.1,463,2.98,1.7)
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

// func TestPositionWithLeapfrog(t *testing.T) {
	// 
// }

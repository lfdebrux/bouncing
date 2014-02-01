package bouncing

import "testing"

func makePosJ(phi0,beta0,v,psi,thetadash float64) *J {
	return &J{P:&P{Phi:phi0,Beta:beta0},V:v,Psi:psi,ThetaDash:thetadash}
}

func BenchmarkPositionJump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ButlerPositionJump(makePosJ(1.35,3.1,463,2.98,1.7))
	}
}

func BenchmarkVondrakPositionJump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VondrakPositionJump(makePosJ(1.35,3.1,463,2.98,1.7))
	}
}

// func TestPositionWithLeapfrog(t *testing.T) {
	// 
// }

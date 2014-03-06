package bouncing

import "testing"

import "math"

func makePosJ(phi0,beta0,v,psi,thetadash float64) *J {
	return &J{P:&P{Phi:phi0,Beta:beta0},Velocity:v,Psi:psi,ThetaDash:thetadash}
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

func TestButlerPositionJumpLongitude(t *testing.T) {
	for v0 := 0.0; v0 < 2000; v0 += 1 {
		j := makePosJ(math.Pi/2,0,v0,math.Pi/2,math.Pi/4)
		ButlerPositionJump(j)
		if j.Beta < 0 {
			t.Fatal("ButlerPositionJump: j.Beta is negative")
		}
		if j. Beta > 2*math.Pi {
			t.Fatal("ButlerPositionJump: j.Beta is greater than 2pi")
		}
	}
}

func TestButlerPositionJumpColatitude(t *testing.T) {
	for v0 := 0.0; v0 < 2000; v0 += 1 {
		j := makePosJ(math.Pi/2,0,v0,0,math.Pi/4)
		ButlerPositionJump(j)
		if j.Phi < 0 {
			t.Fatal("ButlerPositionJump: j.Phi is negative")
		}
		if j. Phi > math.Pi {
			t.Fatal("ButlerPositionJump: j.Phi is greater than pi")
		}
	}
}

func TestVondrakPositionJumpLongitude(t *testing.T) {
	for v0 := 0.0; v0 < 2000; v0 += 1 {
		j := makePosJ(math.Pi/2,0,v0,math.Pi/2,math.Pi/4)
		VondrakPositionJump(j)
		if j.Beta < 0 {
			t.Fatal("VondrakPositionJump: j.Beta is negative")
		}
		if j. Beta > 2*math.Pi {
			t.Fatal("VondrakPositionJump: j.Beta is greater than 2pi")
		}
	}
}

func TestVondrakPositionJumpColatitude(t *testing.T) {
	for v0 := 0.0; v0 < 2000; v0 += 1 {
		j := makePosJ(math.Pi/2,0,v0,0,math.Pi/4)
		VondrakPositionJump(j)
		if j.Phi < 0 {
			t.Fatal("VondrakPositionJump: j.Phi is negative")
		}
		if j. Phi > math.Pi {
			t.Fatal("VondrakPositionJump: j.Phi is greater than pi")
		}
	}
}

func TestToLongitude(t *testing.T) {
	if beta := toLongitude(0); beta != 0 {
		t.Errorf("toLongitude(0) should be 0, instead go %v",beta)
	}
	if beta := toLongitude(math.Pi); beta != math.Pi {
		t.Errorf("toLongitude(pi) should be pi, instead got %vpi",beta/math.Pi)
	}
	if beta := toLongitude(3*math.Pi); beta != math.Pi {
		t.Errorf("toLongitude(3pi) should be pi, instead got %vpi",beta/math.Pi)
	}
	if beta := toLongitude(-math.Pi); beta != math.Pi {
		t.Errorf("toLongitude(-pi) should be pi, instead got %vpi",beta/math.Pi)
	}
	if beta := toLongitude(-math.Pi/2); beta != 3*math.Pi/2 {
		t.Errorf("toLongitude(-pi/2) should be 3pi/2, instead got %vpi",beta/math.Pi)
	}
	if beta := toLongitude(-5*math.Pi/2); beta != 3*math.Pi/2 {
		t.Errorf("toLongitude(-5pi/2) should be 3pi/2, instead got %vpi",beta/math.Pi)
	}
}

// func TestPositionWithLeapfrog(t *testing.T) {
	// 
// }

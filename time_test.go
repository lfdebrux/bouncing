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

func testVondrakZenith(phi, beta float64) (float64, float64) {
	j := &J{P:&P{Beta:beta,Phi:phi}}
	VondrakZenith(j)
	return j.SolarZenith, j.Time
}

func TestVondrakZenithTime(t *testing.T) {
	if _, time := testVondrakZenith(math.Pi/2, 0); time != 0 {
		t.Errorf("VondrakZenith(beta = 0) should have j.Time = 0, instead got %v",time)
	}
	if _, time := testVondrakZenith(math.Pi/2, math.Pi/2); time != math.Pi*TIMEPERRAD {
		t.Errorf("VondrakZenith(beta = pi/2) (dusk terminator) should have j.Time = pi*TIMEPERRAD, instead got %v TIMEPERRAD",time/TIMEPERRAD)
	}
	if _, time := testVondrakZenith(math.Pi/2, math.Pi); time != TIMEPERRAD/2 {
		t.Errorf("VondrakZenith(beta = pi/2) should have j.Time = TIMEPERRAD/2, instead got %v TIMEPERRAD",time/TIMEPERRAD)
	}
	if _, time := testVondrakZenith(math.Pi/2, 3*math.Pi/2); time != 0 {
		t.Errorf("VondrakZenith(beta = 3pi/2) (dawn terminator) should have j.Time = 0, instead got %v TIMEPERRAD",time/TIMEPERRAD)
	}
}
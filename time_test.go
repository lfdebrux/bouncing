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

func testVondrakZenithBeta(t *testing.T, phi, beta, beta1 float64, msg string) {
	j := &J{P:&P{Beta:beta,Phi:phi}}
	VondrakZenith(j)
	if j.Beta != beta1 {
		t.Errorf(msg,j.Beta/math.Pi)
	}
}

func TestVondrakZenithBeta(t *testing.T) {
	testVondrakZenithBeta(t, math.Pi/2, 0, 0, "VondrakZenith(pi/2, 0) -- j.Beta = 0 not %v pi")
	testVondrakZenithBeta(t, 0, 0, 0, "VondrakZenith(0, 0) -- j.Beta = 0 not %v pi")
	testVondrakZenithBeta(t, math.Pi/2, math.Pi/2, math.Pi/2, "VondrakZenith(pi/2, pi/2) -- j.Beta = pi/2 not %v pi")
	testVondrakZenithBeta(t, math.Pi/2, math.Pi/2 + 1e-16, 3*math.Pi/2, "VondrakZenith(pi/2, pi/2 + small) (just after dusk terminator) -- j.Beta = 3pi/2 not %v pi")
	testVondrakZenithBeta(t, math.Pi/2, math.Pi, 3*math.Pi/2, "VondrakZenith(pi/2, pi) -- j.Beta = 3pi/2 not %v pi")
}

func testVondrakZenithSolarZenith(t *testing.T, phi, beta, solarzenith float64, msg string) {
	j := &J{P:&P{Beta:beta,Phi:phi}}
	VondrakZenith(j)
	if j.SolarZenith != solarzenith {
		t.Errorf(msg,j.SolarZenith/math.Pi)
	}
}

func TestVondrakZenithSolarZenith(t *testing.T) {
	testVondrakZenithSolarZenith(t, math.Pi/2, 0, 0, "VondrakZenith(pi/2, 0) -- j.SolarZenith = 0 not %v pi")
	testVondrakZenithSolarZenith(t, math.Pi/2, math.Pi/2, math.Pi/2, "VondrakZenith(pi/2, pi/2) -- j.SolarZenith = pi/2 not %v pi")
	testVondrakZenithSolarZenith(t, math.Pi/2, math.Pi, math.Pi/2, "VondrakZenith(pi/2, pi) -- j.SolarZenith = pi/2 not %v pi")
	testVondrakZenithSolarZenith(t, 0, 0, math.Pi/2, "VondrakZenith(0, 0) -- j.SolarZenith = pi/2 not %v pi")
	testVondrakZenithSolarZenith(t, math.Pi/2, 3*math.Pi/2, math.Pi/2, "VondrakZenith(pi/2, 3pi/2) -- j.SolarZenith = pi/2 not %v pi")
}

func testVondrakZenithTime(t *testing.T, phi, beta, time float64, msg string) {
	j := &J{P:&P{Beta:beta,Phi:phi}}
	VondrakZenith(j)
	if !almosteq(j.Time, time) {
		t.Errorf(msg, j.Time/timeperrad/math.Pi)
	}
}

func TestVondrakZenithTime(t *testing.T) {
	testVondrakZenithTime(t, math.Pi/2, 0, 0, "VondrakZenith(beta = 0) -- j.Time = 0 not %v pi timeperrad")
	testVondrakZenithTime(t, 0, 0, 0, "VondrakZenith(0, 0) -- j.Time = 0 not %v pi timeperrad")
	testVondrakZenithTime(t, math.Pi/2, math.Pi/2, 0, "VondrakZenith(beta = pi/2) (dusk terminator) -- j.Time = 0 not %v pi timeperrad")
	testVondrakZenithTime(t, math.Pi/2, math.Pi/2 + 1e-16, math.Pi*timeperrad, "VondrakZenith(beta = pi/2 + small) (just after dusk terminator) -- j.Time = pi*timeperrad not %v pi timeperrad")
	testVondrakZenithTime(t, math.Pi/2, math.Pi, math.Pi*timeperrad/2, "VondrakZenith(beta = pi) -- j.Time = pi/2 timeperrad not %v pi timeperrad")
	testVondrakZenithTime(t, math.Pi/2, 3*math.Pi/2, 0, "VondrakZenith(beta = 3pi/2) (dawn terminator) -- j.Time = 0 not %v pi timeperrad")
	testVondrakZenithTime(t, math.Pi/2, 3*math.Pi/2 - 1e-19, 0, "VondrakZenith(beta = 3pi/2 - small) (just before dawn terminator) -- j.Time = 0 not %v pi timeperrad")
}

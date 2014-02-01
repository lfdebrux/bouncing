package bouncing

import "testing"

import "math"

func makeJ(v,thetadash float64) *J {
	return &J{V:v,ThetaDash:thetadash}
}

func BenchmarkFlightTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FlightTime(makeJ(300,1.74))
	}
}

func BenchmarkButlerFlightTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ButlerFlightTime(makeJ(300,1.74))
	}
}

func BenchmarkVondrakFlightTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VondrakFlightTime(makeJ(300,1.74))
	}
}

func TestButlerFlightTimeZero(t *testing.T) {
	j := makeJ(0,0)
	ButlerFlightTime(j)
	if j.T != 0 {
		t.Errorf("Expected flight time with v=0 using Butler equation to be 0, instead got %f",j.T)
	}
}

func VaryFlightTime(f JumpMethod, cb func(v,a,t float64)) {
	jump := new(J)
	const NUM = 1000
	for i := 1; i < NUM; i++ {
		for j := 0; j < NUM; j++ {
			jump.V = (2300)*float64(i)/(NUM-1)
			jump.ThetaDash = math.Pi/2*float64(j)/NUM
			f(jump)
			cb(jump.V,jump.ThetaDash,jump.T)
		}
	}
}

func TestFlightTimeIsNum(t *testing.T) {
	VaryFlightTime(FlightTime, func(v,a,f float64) {
		if math.IsNaN(f) || math.IsInf(f,0) {
			t.Fatalf("FlightTime(%f,%f) = %v, not a number",v,a,f)
		}
	})
}

func TestButlerFlightTimeIsNum(t *testing.T) {
	VaryFlightTime(ButlerFlightTime, func(v,a,f float64) {
		if math.IsNaN(f) || math.IsInf(f,0) {
			t.Fatalf("ButlerFlightTime(%f,%f) = %v, not a number",v,a,f)
		}
	})
}

func TestVondrakFlightTimeIsNum(t *testing.T) {
	VaryFlightTime(VondrakFlightTime, func(v,a,f float64) {
		if math.IsNaN(f) || math.IsInf(f,0) {
			t.Fatalf("VondrakFlightTime(%f,%f) = %v, not a number",v,a,f)
		}
	})
}

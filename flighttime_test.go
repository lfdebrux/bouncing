package bouncing

import "testing"

import "math"

func BenchmarkFlightTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FlightTime(300,1.74)
	}
}

func BenchmarkButlerFlightTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ButlerFlightTime(300,1.74)
	}
}

func BenchmarkVondrakFlightTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VondrakFlightTime(300,1.74)
	}
}

func TestButlerFlightTimeZero(t *testing.T) {
	if time := ButlerFlightTime(0,0); time != 0 {
		t.Errorf("Expected flight time with v=0 using Butler equation to be 0, instead got %f",time)
	}
}

func VaryFlightTime(f F, cb func(v,a,t float64)) {
	const NUM = 1000
	for i := 1; i < NUM; i++ {
		for j := 0; j < NUM; j++ {
			v := (2300)*float64(i)/(NUM-1)
			a := math.Pi/2*float64(j)/NUM
			t := f(v,a)
			cb(v,a,t)
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

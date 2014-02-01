package bouncing

import "testing"

import "math"
import "math/rand"

func makeTempJ(phi float64) *J {
	return &J{P:&P{Phi:phi}}
}

func testT(f JumpMethod,phi float64) float64 {
	j := makeTempJ(phi)
	f(j)
	return j.Temperature
} 

func BenchmarkTemperature(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ButlerTemperature(makeTempJ(rand.Float64()*math.Pi))
	}
}

func TestTemperaturePoleLimit(t *testing.T) {
	if T := testT(ButlerTemperature,0);math.Abs(T - ButlerT0) > TOL {
		t.Log(math.Abs(T - ButlerT0))
		t.Errorf("Temperature at Lunar Pole should be ButlerT0=%f, instead got %f",ButlerT0,T)
	}
}

func TestTemperatureEquatorLimit(t *testing.T) {
	Te := ButlerT0 + ButlerT1
	if T := testT(ButlerTemperature,math.Pi/2); math.Abs(T - Te) > TOL {
		t.Log(math.Abs(T - Te))
		t.Errorf("Temperature at Lunar equator should be ButlerT0+ButlerT1=%f, instead got %f",Te,T)
	}
}

func TestTemperatureButlerT0IsMin(t *testing.T) {
	for phi := 0.0; phi < math.Pi/2; phi += 0.1 {
		if T := testT(ButlerTemperature,phi); T < ButlerT0 {
			t.Errorf("Temperature should never be lower than ButlerT0=%f, but at phi %f got %f",ButlerT0,phi,T)
		}
	}
}

func TestTemperatureTeIsMax(t *testing.T) {
	Te := ButlerT0 + ButlerT1
	for phi := 0.0; phi < math.Pi/2; phi += 0.1 {
		if T := testT(ButlerTemperature,phi); T > Te {
			t.Errorf("Temperature should never be higher than Te=%f, but at phi %f got %f",Te,phi,T)
		}
	}
}

func TestTemperatureMonotonic(t *testing.T) {
	Tlast := testT(ButlerTemperature,0)
	for phi := 0.0; phi < math.Pi/2; phi += 0.1 {
		if T := testT(ButlerTemperature,phi); T < Tlast {
			t.Errorf("Temperature should vary monotonically, but it didn't")
		}
	}
}

func TestTemperatureSymmetric(t *testing.T) {
	for phi := 0.0; phi < math.Pi/2; phi += 0.1 {
		if math.Abs(testT(ButlerTemperature,phi) - testT(ButlerTemperature,math.Pi-phi)) > TOL {
			t.Log(testT(ButlerTemperature,phi) - testT(ButlerTemperature,math.Pi-phi))
			t.Errorf("Temperature function should be symmetric, but Temperature(%f)=%f!=Temperature(%f)=%f",phi,testT(ButlerTemperature,phi),math.Pi-phi,testT(ButlerTemperature,math.Pi-phi))
		}
	}
}

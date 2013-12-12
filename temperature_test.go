package bouncing

import "testing"

import "math"
import "math/rand"

func BenchmarkTemperature(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ButlerTemperature(rand.Float64()*math.Pi,0)
	}
}

func TestTemperaturePoleLimit(t *testing.T) {
	if T := ButlerTemperature(0,0); math.Abs(T - ButlerT0) > TOL {
		t.Log(math.Abs(T - ButlerT0))
		t.Errorf("Temperature at Lunar Pole should be ButlerT0=%f, instead got %f",ButlerT0,T)
	}
}

func TestTemperatureEquatorLimit(t *testing.T) {
	Te := ButlerT0 + ButlerT1
	if T := ButlerTemperature(math.Pi/2,0); math.Abs(T - Te) > TOL {
		t.Log(math.Abs(T - Te))
		t.Errorf("Temperature at Lunar equator should be ButlerT0+ButlerT1=%f, instead got %f",Te,T)
	}
}

func TestTemperatureButlerT0IsMin(t *testing.T) {
	for phi := 0.0; phi < math.Pi/2; phi += 0.1 {
		if T := ButlerTemperature(phi,0); T < ButlerT0 {
			t.Errorf("Temperature should never be lower than ButlerT0=%f, but at phi %f got %f",ButlerT0,phi,T)
		}
	}
}

func TestTemperatureTeIsMax(t *testing.T) {
	Te := ButlerT0 + ButlerT1
	for phi := 0.0; phi < math.Pi/2; phi += 0.1 {
		if T := ButlerTemperature(phi,0); T > Te {
			t.Errorf("Temperature should never be higher than Te=%f, but at phi %f got %f",Te,phi,T)
		}
	}
}

func TestTemperatureMonotonic(t *testing.T) {
	Tlast := ButlerTemperature(0,0)
	for phi := 0.0; phi < math.Pi/2; phi += 0.1 {
		if T := ButlerTemperature(phi,0); T < Tlast {
			t.Errorf("Temperature should vary monotonically, but it didn't")
		}
	}
}

func TestTemperatureSymmetric(t *testing.T) {
	for phi := 0.0; phi < math.Pi/2; phi += 0.1 {
		if math.Abs(ButlerTemperature(phi,0) - ButlerTemperature(math.Pi-phi,0)) > TOL {
			t.Log(ButlerTemperature(phi,0) - ButlerTemperature(math.Pi-phi,0))
			t.Errorf("Temperature function should be symmetric, but Temperature(%f)=%f!=Temperature(%f)=%f",phi,ButlerTemperature(phi,0),math.Pi-phi,ButlerTemperature(math.Pi-phi,0))
		}
	}
}

package bouncing

import "testing"

import "math"

func TestTemperaturePoleLimit(t *testing.T) {
	if T := Temperature(0); math.Abs(T - T0) > TOL {
		t.Log(math.Abs(T - T0))
		t.Errorf("Temperature at Lunar Pole should be T0=%f, instead got %f",T0,T)
	}
}

func TestTemperatureEquatorLimit(t *testing.T) {
	Te := T0 + T1
	if T := Temperature(math.Pi/2); math.Abs(T - Te) > TOL {
		t.Log(math.Abs(T - Te))
		t.Errorf("Temperature at Lunar equator should be T0+T1=%f, instead got %f",Te,T)
	}
}

func TestTemperatureT0IsMin(t *testing.T) {
	for phi := 0.0; phi < math.Pi/2; phi += 0.1 {
		if T := Temperature(phi); T < T0 {
			t.Errorf("Temperature should never be lower than T0=%f, but at phi %f got %f",T0,phi,T)
		}
	}
}

func TestTemperatureTeIsMax(t *testing.T) {
	Te := T0 + T1
	for phi := 0.0; phi < math.Pi/2; phi += 0.1 {
		if T := Temperature(phi); T > Te {
			t.Errorf("Temperature should never be higher than Te=%f, but at phi %f got %f",Te,phi,T)
		}
	}
}

func TestTemperatureMonotonic(t *testing.T) {
	Tlast := Temperature(0)
	for phi := 0.0; phi < math.Pi/2; phi += 0.1 {
		if T := Temperature(phi); T < Tlast {
			t.Errorf("Temperature should vary monotonically, but it didn't")
		}
	}
}

func TestTemperatureSymmetric(t *testing.T) {
	for phi := 0.0; phi < math.Pi/2; phi += 0.1 {
		if math.Abs(Temperature(phi) - Temperature(math.Pi-phi)) > TOL {
			t.Log(Temperature(phi) - Temperature(math.Pi-phi))
			t.Errorf("Temperature function should be symmetric, but Temperature(%f)=%f!=Temperature(%f)=%f",phi,Temperature(phi),math.Pi-phi,Temperature(math.Pi-phi))
		}
	}
}
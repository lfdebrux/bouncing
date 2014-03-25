package bouncing

import "testing"

import "math"

func BenchmarkIsLost(b *testing.B) {
	j := newJ()
	for i := 0; i < b.N; i++ {
		IsLost(j)
	}
}

func percentPhoto(j *J) float64 {
	j.Phi = math.Pi/2
	j.Velocity = 0
	var count float64
	for i := 0; i < NUM; i++ {
		if l := IsLost(j); l != nil {
			if l.HowLost == Photodestruction {
				count++
			} else {
				panic("loss_test: percentPhoto: should only be able to lose particles to Photodestruction")
			}
		}
	}
	return count/NUM
}

func TestIsPhotoZero(t *testing.T) {
	j := &J{P:new(P)}

	j.FlightTime = 0

	if pc := percentPhoto(j); pc != 0 {
		t.Errorf("No particles should be photodestructed when flight time is 0, however %f%% were lost",pc)
	}
}

func TestIsPhotoTau(t *testing.T) {
	j := &J{P:new(P)}

	j.FlightTime = Tau

	if pc := percentPhoto(j); !almosteq(pc, 1 - 1/math.E) {
		t.Log(diff(pc,1-1/math.E))
		t.Errorf("Expect 1 - 1/e=%f particles to be photodestructed for flight time tau=%f, instead %f%% were lost",1-1/math.E,Tau,pc)
	}
}

func percentCapture(j *J) float64 {
	var count float64
	for i := 0; i < NUM; i++ {
		if l := IsCaptureButler(j); l != nil {
			count++
		}
	}
	return count/NUM
}

func TestIsCaptureEquator(t *testing.T) {
	j := &J{P:new(P)}

	j.Phi = math.Pi/2

	if pc := percentCapture(j); pc != 0 {
		t.Errorf("No particles should be captured at equator, however %f%% were lost",pc)
	}
}

func testButlerFstable(lat float64) float64 {
	return fstable[int(math.Ceil(lat*18/math.Pi))-6]
}

func TestButlerFstableLessThan60(t *testing.T) {
	if g := testButlerFstable(5.1*math.Pi/18); g != fstable[0] {
		t.Errorf("ButlerFstable at 5*Pi/18 < latitude < 6*Pi/18 should be %f, instead got %f",fstable[0],g)
	}
}

func TestButlerFstable(t *testing.T) {
	for i,f := range fstable {
		l := (5.5+float64(i))*math.Pi/18
		if g := testButlerFstable(l); g != f {
			t.Log(l*18/math.Pi)
			t.Errorf("ButlerFstable at %d*Pi/18 latitude should be %f, instead got %f",6+i,f,g)
		}
	}
}

func TestIsCaptureExpectedValues(t *testing.T) {
	for i,f := range fstable {
		j := &J{P:new(P)}
		j.Phi = math.Pi/2 - (6+float64(i))*math.Pi/18
		if pc := percentCapture(j); !almosteq(pc,f) {
			t.Log(diff(pc,f))
			t.Errorf("Expect to capture %f%% at %d*Pi/18 latitude (%f), instead lost %f",f*100,6+i,j.Phi,pc*100)
		}
	}
}

func BenchmarkIsNaN(b *testing.B) {
	j := &J{P:new(P)}
	j.Phi = math.NaN()
	for i := 0; i < b.N; i++ {
		IsNaN(j)
	}
}

func TestIsNaN(t *testing.T) {
	j := &J{P:new(P)}
	j.Phi = math.NaN()
	loss := IsNaN(j)
	if loss == nil || loss.HowLost != Error {
		t.Errorf("IsNaN did not detect NaN, loss=%v", loss)
	} else if loss.Error() != "loss: j.Phi is NaN" {
		t.Errorf("IsNaN did not detect correct NaN, loss=%v", loss)
	}
}

func TestIsNaNMultiple(t *testing.T) {
	j := &J{P:new(P)}
	j.Beta = math.NaN()
	j.Temperature = math.NaN()

	loss := IsNaN(j)
	if loss == nil || loss.Error() != "loss: j.Beta is NaN j.Temperature is NaN" {
		t.Errorf("IsNaN did not correctly detect multiple NaN, loss=%v",loss)
	}
}

var names = map[LostType]string{Error:"Error", ThermalEscape:"ThermalEscape", Photodestruction:"Photodestruction", Capture:"Capture"}

func TestThermalEscape(t *testing.T) {
	j := &J{P:new(P), Velocity: 2*Vesc}

	loss := IsLost(j)
	t.Log(loss)
	if loss.HowLost != ThermalEscape {
		t.Errorf("Expecting j.HowLost = %s not ThermalEscape", names[loss.HowLost])
	}
}

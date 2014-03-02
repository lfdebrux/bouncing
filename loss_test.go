package bouncing

import "testing"

import "math"

func percentPhoto(j *J) float64 {
	j.Phi = math.Pi/2
	j.Velocity = 0
	var count float64
	for i := 0; i < NUM; i++ {
		if l := j.IsLost(); l != nil {
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

	j.FlightTime = TAU

	if pc := percentPhoto(j); !almosteq(pc, 1 - 1/math.E) {
		t.Log(diff(pc,1-1/math.E))
		t.Errorf("Expect 1 - 1/e=%f particles to be photodestructed for flight time tau=%f, instead %f%% were lost",1-1/math.E,TAU,pc)
	}
}

func percentCapture(j *J) float64 {
	var count float64
	for i := 0; i < NUM; i++ {
		if l := j.IsCapture(); l != nil {
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

func TestFstableLessThan60(t *testing.T) {
	if g := fstable(5.1*math.Pi/18); g != FSTABLE[0] {
		t.Errorf("fstable at 5*Pi/18 < latitude < 6*Pi/18 should be %f, instead got %f",FSTABLE[0],g)
	}
}

func TestFstable(t *testing.T) {
	for i,f := range FSTABLE {
		l := (5.5+float64(i))*math.Pi/18
		if g := fstable(l); g != f {
			t.Log(l*18/math.Pi)
			t.Errorf("fstable at %d*Pi/18 latitude should be %f, instead got %f",6+i,f,g)
		}
	}
}

func TestIsCaptureExpectedValues(t *testing.T) {
	for i,f := range FSTABLE {
		j := &J{P:new(P)}
		j.Phi = math.Pi/2 - (6+float64(i))*math.Pi/18
		if pc := percentCapture(j); !almosteq(pc,f) {
			t.Log(diff(pc,f))
			t.Errorf("Expect to capture %f%% at %d*Pi/18 latitude (%f), instead lost %f",f*100,6+i,j.Phi,pc*100)
		}
	}
}
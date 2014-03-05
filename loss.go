package bouncing

import "fmt"
import "math"
import "math/rand"

var TAU = 6.7e4 // seconds // from Butler1997
var FSTABLE = []float64{0.4e-2,0.9e-2,4e-2,11e-2}

type LostType int

const (
	Error LostType = iota
	ThermalEscape
	Photodestruction
	Capture
)

type Lost struct {
	err string
	HowLost LostType
}

func (l *Lost) Error() string {
	return l.err
}

func IsNaN(j *J) *Lost {
	msg := "loss:"
	if math.IsNaN(j.Phi) {
		msg += " j.Phi is NaN"
	}
	if math.IsNaN(j.Beta) {
		msg += " j.Beta is NaN"
	}
	if math.IsNaN(j.Time) {
		msg += " j.Time is NaN"
	}
	if math.IsNaN(j.SolarZenith) {
		msg += " j.SolarZenith is NaN"
	}
	if math.IsNaN(j.Velocity) {
		msg += " j.Velocity is NaN"
	}
	if math.IsNaN(j.Psi) {
		msg += " j.Psi is NaN"
	}
	if math.IsNaN(j.ThetaDash) {
		msg += " j.ThetaDash is NaN"
	}
	if math.IsNaN(j.Temperature) {
		msg += " j.Temperature is NaN"
	}
	if math.IsNaN(j.FlightTime) {
		msg += " j.FlightTime is NaN"
	}
	// no NaNs detected
	if msg == "loss:" {
		return nil
	}
	return &Lost{msg, Error}
}

func IsLost(j *J) *Lost {
	if j.Velocity > VESC {
		return &Lost{fmt.Sprintf("loss: thermal escape, v=%f",j.Velocity),ThermalEscape}
	}
	if l := IsNaN(j); l != nil {
		return l
	}
	if rand.Float64() > math.Exp(-j.FlightTime/TAU) {
		return &Lost{fmt.Sprintf("loss: photodestruction, t=%f",j.FlightTime),Photodestruction}
	}
	return nil
}

func IsCaptureButler(j *J) *Lost {
	lat := math.Abs(math.Pi/2 - j.Phi)
	if lat > 5*math.Pi/18 {
		if rand.Float64() < FSTABLE[int(math.Ceil(lat*18/math.Pi))-6] {
			return &Lost{fmt.Sprintf("loss: capture by stable region, phi=%f",j.Phi),Capture}
		}
	}
	return nil
}

func IsCaptureVondrak(j *J) *Lost {
	lat := j.Phi - math.Pi/2
	var fstable float64
	if lat > 85*math.Pi/180 {
		fstable = 0.0367
	} else if lat < -85*math.Pi/180 {
		fstable = 0.0706
	} else {
		return nil
	}
	if rand.Float64() < fstable {
		return &Lost{fmt.Sprintf("loss: capture by stable region, phi=%f",j.Phi),Capture}
	}
	return nil
}

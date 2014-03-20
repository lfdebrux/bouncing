package bouncing

import (
	"fmt"

	"math"
	"math/rand"
)

var Tau = 6.7e4 // seconds // from Butler1997

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

func NewLost(msg string, typ LostType) *Lost {
	return &Lost{err:msg,HowLost:typ}
}

// TODO: reduce repition here
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
		return &Lost{fmt.Sprintf("loss: %v thermal escape, v=%f", j.Time, j.Velocity), ThermalEscape}
	}
	if l := IsNaN(j); l != nil {
		return l
	}
	if rand.Float64() > math.Exp(-j.FlightTime/Tau) {
		return &Lost{fmt.Sprintf("loss: %v photodestruction, flighttime=%f", j.Time, j.FlightTime), Photodestruction}
	}
	return nil
}

var fstable = []float64{0.4e-2,0.9e-2,4e-2,11e-2}
func IsCaptureButler(j *J) *Lost {
	lat := math.Abs(math.Pi/2 - j.Phi)
	if lat > 5*math.Pi/18 {
		if rand.Float64() < fstable[int(math.Ceil(lat*18/math.Pi))-6] {
			return &Lost{fmt.Sprintf("loss: %v capture by stable region, (beta, phi)=(%v, %v)", j.Time, j.Beta, j.Phi), Capture}
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
			return &Lost{fmt.Sprintf("loss: %v capture by stable region, (beta, phi)=(%v, %v)", j.Time, j.Beta, j.Phi), Capture}
	}
	return nil
}

// TODO: reduce repetition here
func SanityChecks(j *J) *Lost {
	msg := "error:"
	if j.Time < 0 {
		msg += " j.Time < 0"
	}
	if j.FlightTime < 0 {
		msg += " j.FlightTime < 0"
	}
	if j.Velocity < 0 {
		msg += " j.Velocity < 0"
	}
	if j.Temperature < 0 {
		msg += " j.Temperature < 0"
	}
	if j.Phi < 0 {
		msg += " j.Phi < 0"
	}
	if j.Phi > math.Pi {
		msg += " j.Phi > pi"
	}
	if j.Beta < 0 {
		msg += " j.Beta < 0"
	}
	if j.Beta > 2*math.Pi {
		msg += " j. Beta > 2pi"
	}
	if j.Psi < 0 {
		msg += " j.Psi < 0"
	}
	if j.Psi > 2*math.Pi {
		msg += " j.Psi > 2pi"
	}
	if j.ThetaDash < 0 {
		msg += " j.ThetaDash < 0"
	}
	if j.ThetaDash > math.Pi/2 {
		msg += " j.ThetaDash > pi/2"
	}
	if j.SolarZenith < 0 {
		msg += " j.SolarZenith < 0"
	}
	// all tests passed
	if msg == "error:" {
		return nil
	}
	return NewLost(msg,Error)
}

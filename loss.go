package bouncing

import (
	"fmt"

	"math"
	"math/rand"
)

var Tau = 6.7e4 // seconds // from Butler1997

func IsLost(j *J) bool {
	if j.Type > isLost {
		return true
	}
	return false
}

func NewLost(j *J, typ JumpType) string {
	j.Type = typ
	switch typ {
	case ThermalEscape:
		return fmt.Sprintf("l %v thermal escape, v=%f", j.Time, j.Velocity)
	case Photodestruction:
		return fmt.Sprintf("l %v photodestruction, flighttime=%f", j.Time, j.FlightTime)
	case Capture:
		return fmt.Sprintf("l %v capture by stable region, (beta, phi)=(%v, %v)", j.Time, j.Beta, j.Phi)
	default:
		return fmt.Sprintf("l %v", j.Time)
	}
}

var fstable = []float64{0.4e-2,0.9e-2,4e-2,11e-2}
func CaptureButler(j *J) string {
	lat := math.Abs(math.Pi/2 - j.Phi)
	if lat > 5*math.Pi/18 {
		if rand.Float64() < fstable[int(math.Ceil(lat*18/math.Pi))-6] {
			return NewLost(j, Capture)
		}
	}
	return ""
}

func CaptureVondrak(j *J) string {
	lat := j.Phi - math.Pi/2
	var fstable float64
	if lat > 85*math.Pi/180 {
		fstable = 0.0367
	} else if lat < -85*math.Pi/180 {
		fstable = 0.0706
	} else {
		return ""
	}
	if rand.Float64() < fstable {
			return NewLost(j, Capture)
	}
	return ""
}

func CheckLost(j *J) string {
	if j.Velocity > Vesc {
		return NewLost(j, ThermalEscape)
	}
	if err := CheckNaN(j); err != "" {
		return err
	}
	if rand.Float64() > math.Exp(-j.FlightTime/Tau) {
		return NewLost(j, Photodestruction)
	}
	return ""
}

// TODO: reduce repetition here
func CheckNaN(j *J) string {
	msg := "e"
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
	if msg == "e" {
		return ""
	}
	j.Type = Error
	return msg
}

// TODO: reduce repetition here
func CheckSanity(j *J) string {
	msg := "e"
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
	if msg == "e" {
		return ""
	}
	j.Type = Error
	return msg
}

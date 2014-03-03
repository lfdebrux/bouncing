package bouncing

import "fmt"
import "math"
import "math/rand"

const TAU = 6.7e4 // seconds
var FSTABLE = []float64{0.4e-2,0.9e-2,4e-2,11e-2}

type LostType int

const (
	ThermalEscape LostType = iota
	Photodestruction
	Capture
)

type Lost struct {
	HowLost LostType
	Jump *J
}

func (l *Lost) Error() string {
	switch l.HowLost {
		case ThermalEscape:
			return fmt.Sprintf("loss: particle lost due to thermal escape, v=%f",l.Jump.Velocity)
		case Photodestruction:
			return fmt.Sprintf("loss: particle lost due to photodestruction, t=%f",l.Jump.FlightTime)
		case Capture:
			return fmt.Sprintf("loss: particle lost due to capture by stable region, phi=%f",l.Jump.Phi)
	}
	return "particle lost"
}

func (j *J) IsLost() *Lost {
	if j.Velocity > VESC {
		return &Lost{ThermalEscape,j}
	}
	if rand.Float64() > math.Exp(-j.FlightTime/TAU) {
		return &Lost{Photodestruction,j}
	}
	if l := j.IsCapture(); l != nil {
		return l
	}
	return nil
}

func fstable(lat float64) float64 {
	return FSTABLE[int(math.Ceil(lat*18/math.Pi))-6]	
}

func (j *J) IsCapture() *Lost {
	lat := math.Abs(math.Pi/2 - j.Phi)
	if lat > 5*math.Pi/18 {
		if rand.Float64() < fstable(lat) {
			return &Lost{Capture,j}
		}
	}
	return nil
}

func VondrakColdTraps(j *J) {
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
		return &Lost{Capture,j}
	}
	return nil
}

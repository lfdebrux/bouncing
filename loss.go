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
			return fmt.Sprintf("particle lost due to thermal escape, v=%f",l.Jump.V)
		case Photodestruction:
			return fmt.Sprintf("particle lost due to photodestruction, t=%f",l.Jump.T)
		case Capture:
			return fmt.Sprintf("particle lost due to capture by stable region, phi=%f",l.Jump.Phi)
	}
	return "particle lost"
}

func (j *J) IsLost() *Lost {
	if j.V > VESC {
		return &Lost{ThermalEscape,j}
	}
	if rand.Float64() > math.Exp(-j.T/TAU) {
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

package bouncing

import "fmt"
import "math"
import "math/rand"

const TAU = 6.7e4 // seconds
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
	Jump *J
}

func (l *Lost) Error() string {
	return l.err
}

func (j *J) NaNCheck() *Lost {
	switch {
	case math.IsNaN(j.Phi):
		return &Lost{"loss: j.Phi is NaN",Error,j}
	case math.IsNaN(j.Beta):
		return &Lost{"loss: j.Beta is NaN",Error,j}
	case math.IsNaN(j.Time):
		return &Lost{"loss: j.Time is NaN",Error,j}
	case math.IsNaN(j.SolarZenith):
		return &Lost{"loss: j.SolarZenith is NaN",Error,j}
	case math.IsNaN(j.Velocity):
		return &Lost{"loss: j.Velocity is NaN",Error,j}
	case math.IsNaN(j.Psi):
		return &Lost{"loss: j.Psi is NaN",Error,j}
	case math.IsNaN(j.ThetaDash):
		return &Lost{"loss: j.ThetaDash is NaN",Error,j}
	case math.IsNaN(j.Temperature):
		return &Lost{"loss: j.Temperature is NaN",Error,j}
	case math.IsNaN(j.FlightTime):
		return &Lost{"loss: j.FlightTime is NaN",Error,j}
	}
	return nil
}

func (j *J) IsLost() *Lost {
	if j.Velocity > VESC {
		return &Lost{fmt.Sprintf("loss: particle lost due to thermal escape, v=%f",j.Velocity),ThermalEscape,j}
	}
	if rand.Float64() > math.Exp(-j.FlightTime/TAU) {
		return &Lost{fmt.Sprintf("loss: particle lost due to photodestruction, t=%f",j.FlightTime),Photodestruction,j}
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
			return &Lost{fmt.Sprintf("loss: particle lost due to capture by stable region, phi=%f",j.Phi),Capture,j}
		}
	}
	return nil
}

func VondrakColdTraps(j *J) *Lost {
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
		return &Lost{fmt.Sprintf("loss: particle lost due to capture by stable region, phi=%f",j.Phi),Capture,j}
	}
	return nil
}

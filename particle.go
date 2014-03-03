package bouncing

import "math"
import "math/rand"

type P struct {
	Type ParticleType
	Phi,Beta float64 // colatitude, longitude radians
	Time float64 // lunar time of day, seconds
	SolarZenith float64 // sun angle from zenith in radians
}

type ParticleType int

const (
	Water = iota
	Hydrogen
)

const AMU = 1.660538921e-27 // kg

var Mass = []float64{
	18*AMU, // Water
	1*AMU, // Hydrogen
}

func RandParticle(typ ParticleType) *P {
	phi := math.Acos(2*rand.Float64()-1)
	beta := 2*math.Pi*rand.Float64()
	return &P{Type:typ,Phi:phi,Beta:beta}
}

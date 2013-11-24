package bouncing

import "math"
import "math/rand"

type P struct {
	Type ParticleType
	Phi,Beta float64
}

type ParticleType int

const (
	Water = iota
)

const AMU = 1.660538921e-27 // kg

var Mass = []float64{
	18*AMU, // Water
}

func RandParticle(typ ParticleType) *P {
	phi := math.Acos(2*rand.Float64()-1)
	beta := 2*math.Pi*rand.Float64()
	return &P{Type:typ,Phi:phi,Beta:beta}
}

package bouncing

type P struct {
	Type ParticleType
	Phi,Beta float64 // colatitude, longitude radians
	Time float64 // lunar time of day, seconds
	SolarZenith float64 // sun angle from zenith in radians
}

type ParticleGenerator func() *P

type ParticleType int

const (
	Water ParticleType = iota
	Hydrogen
)

const amu = 1.660538921e-27 // kg

var Mass = map[ParticleType]float64{
	Water: 18*amu,
	Hydrogen: 1*amu,
}

package bouncing

import (
	"math"
	"math/rand"
)

func RandInitialPositionButler(p *P) {
	p.Phi = math.Acos(2*rand.Float64()-1)
	p.Beta = 2*math.Pi*rand.Float64()
}

func RandInitialPositionVondrak(p *P) {
	p1, p2 := rand.Float64(), 2*math.Pi*rand.Float64()
	p.Phi = math.Acos( math.Sqrt(p1)*math.Sin(p2) )
	beta := math.Asin( math.Sqrt(p1)*math.Cos(p2)/math.Sqrt(1-p1*math.Sin(p2)*math.Sin(p2)) )
	p.Beta = toLongitude(beta)
}

func RandVelocity(j *J) string {
	m,temp := Mass[j.P.Type],j.Temperature
	if Q == nil {
		InitMaxwellian()
	}
	p := rand.Float64()
	a := math.Sqrt(m/(K_B*temp))

	j.Velocity = Q.Eval(p)/a

	if j.Velocity > Vesc {
		return NewLost(j, ThermalEscape)
	}

	return ""
}

func RandDirection(j *J) {
	j.Psi = 2*math.Pi*rand.Float64()
	j.ThetaDash = math.Acos(rand.Float64())
}

func ButlerRandDirection(j *J) {
	j.Psi = 2*math.Pi*rand.Float64()
	j.ThetaDash = rand.Float64()*math.Pi/2
}

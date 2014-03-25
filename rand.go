package bouncing

import (
	"fmt"
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
	p.Beta = math.Asin( math.Sqrt(p1)*math.Cos(p2)/math.Sqrt(1-p1*math.Sin(p2)*math.Sin(p2)) )
}

func RandVelocity(j *J) *Lost {
	m,temp := Mass[j.Type],j.Temperature
	if Q == nil {
		InitMaxwellian()
	}
	p := rand.Float64()
	a := math.Sqrt(m/(K_B*temp))

	j.Velocity = Q.Eval(p)/a

	if j.Velocity > Vesc {
		return &Lost{fmt.Sprintf("loss: thermal escape, v=%f",j.Velocity),ThermalEscape}
	}

	return nil
}

func RandDirection(j *J) *Lost {
	j.Psi = 2*math.Pi*rand.Float64()
	j.ThetaDash = math.Acos(rand.Float64())
	return nil
}

func ButlerRandDirection(j *J) *Lost {
	j.Psi = 2*math.Pi*rand.Float64()
	j.ThetaDash = rand.Float64()*math.Pi/2
	return nil
}

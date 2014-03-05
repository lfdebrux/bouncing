package bouncing

import (
	"fmt"
	"math"
	"math/rand"
)

func RandVelocity(j *J) *Lost {
	m,temp := Mass[j.Type],j.Temperature
	if Q == nil {
		InitMaxwellian()
	}
	p := rand.Float64()
	a := math.Sqrt(m/(K_B*temp))

	j.Velocity = Q.Eval(p)/a

	if j.Velocity > VESC {
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

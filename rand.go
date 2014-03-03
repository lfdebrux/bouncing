package bouncing

import "math"
import "math/rand"

func RandVelocity(j *J) *Lost {
	m,temp := Mass[j.Type],j.Temperature
	if Q == nil {
		InitMaxwellian()
	}
	p := rand.Float64()
	a := math.Sqrt(m/(K_B*temp))

	j.Velocity = Q.Eval(p)/a
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

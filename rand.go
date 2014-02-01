package bouncing

import "math"
import "math/rand"

func RandVelocity(j *J) {
	m,temp := Mass[j.Type],j.Temperature
	if Q == nil {
		InitMaxwellian()
	}
	p := rand.Float64()
	a := math.Sqrt(m/(K_B*temp))

	j.V = Q.Eval(p)/a
}

func RandDirection(j *J) {
	j.Psi = 2*math.Pi*rand.Float64()
	j.ThetaDash = math.Acos(rand.Float64())
}

func ButlerRandDirection(j *J) {
	j.Psi = 2*math.Pi*rand.Float64()
	j.ThetaDash = rand.Float64()*math.Pi/2
}

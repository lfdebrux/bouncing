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

type RandDirectionFunc func() (psi,thetadash float64)

func RandDirection() (psi,thetadash float64) {
	psi = 2*math.Pi*rand.Float64()
	thetadash = math.Acos(rand.Float64())
	return psi,thetadash
}

func ButlerRandDirection() (psi,thetadash float64) {
	psi = 2*math.Pi*rand.Float64()
	thetadash = rand.Float64()*math.Pi/2
	return psi,thetadash	
}

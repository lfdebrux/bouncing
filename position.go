package bouncing

import "math"
import "math/rand"

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

func PositionJump(phi0,beta0,v,psi,thetadash float64) (phi,beta float64) {
	d := 2*math.Atan( 1 /( (VESC*VESC/(v*v))/(2*math.Sin(thetadash)*math.Cos(thetadash)) - math.Tan(thetadash) ) )
	e := 0.0
	if phi0 == 0 {
		phi = d
		e = psi
	} else {
		phi = math.Acos(math.Cos(phi0)*math.Cos(d) + math.Sin(phi0)*math.Sin(d)*math.Cos(psi))
		e = math.Acos((math.Cos(d) - math.Cos(phi)*math.Cos(phi0))/(math.Sin(phi)*math.Sin(phi0)))
	}
	if psi > math.Pi {
		beta = beta0 + e
	} else {
		beta = beta0 - e
	}
	return phi,beta
}

func VondrakPositionJump(phi0,beta0,v,psi,thetadash float64) (phi,beta float64) {
	d := 2*math.Atan(math.Sin(2*thetadash)/((VESC/v)*(VESC/v)-1-math.Cos(2*thetadash)))
	phi = math.Acos(math.Cos(phi0)*math.Cos(d) + math.Sin(phi0)*math.Sin(d)*math.Cos(psi))
	e := math.Acos((math.Cos(d)-math.Cos(phi0)*math.Cos(phi))/(math.Sin(phi0)*math.Sin(phi)))
	if psi > math.Pi {
		beta = beta0 + e
	} else {
		beta = beta0 - e
	}
	return phi,beta
}

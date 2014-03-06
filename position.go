package bouncing

import "math"

func toLongitude(beta float64) float64 {
	beta = math.Mod(beta,2*math.Pi)
	if beta < 0 {
		beta = 2*math.Pi + beta
	}
	return beta
}

func ButlerPositionJump(j *J) *Lost {
	phi0,beta0 := j.Phi,j.Beta
	v,psi,thetadash := j.Velocity,j.Psi,j.ThetaDash

	d := 2*math.Atan( 1 /( (VESC*VESC/(v*v))/(2*math.Sin(thetadash)*math.Cos(thetadash)) - math.Tan(thetadash) ) )

	var e,phi,beta float64
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
	
	j.Phi = phi
	j.Beta = toLongitude(beta)

	return nil
}

func VondrakPositionJump(j *J) *Lost {
	phi0,beta0 := j.Phi,j.Beta
	v,psi,thetadash := j.Velocity,j.Psi,j.ThetaDash

	d := 2*math.Atan(math.Sin(2*thetadash)/((VESC/v)*(VESC/v)-1-math.Cos(2*thetadash)))
	phi := math.Acos(math.Cos(phi0)*math.Cos(d) + math.Sin(phi0)*math.Sin(d)*math.Cos(psi))
	e := math.Acos((math.Cos(d)-math.Cos(phi0)*math.Cos(phi))/(math.Sin(phi0)*math.Sin(phi)))

	var beta float64

	if psi > math.Pi {
		beta = beta0 + e
	} else {
		beta = beta0 - e
	}
	
	j.Phi = phi
	j.Beta = toLongitude(beta)

	return nil
}

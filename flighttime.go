package bouncing

import "math"

func flightTimeMsg(j *J, flighttime float64) string {
	fmt.Println("j", j.Time, j.Phi, j.Beta, flighttime)
}

func FlightTime(j *J) string {
	v,thetadash := j.Velocity,j.ThetaDash
	if v == 0 {
		return
	}
	if v > Vesc {
		return
	}
	b := v*v*R/Mu
	g := 2 - b
	o := math.Sin(thetadash)*math.Sin(thetadash)
	e := math.Sqrt(1-b*g*o)
	T := 2*math.Pi*math.Sqrt(R*R*R/(Mu*g*g*g))

	t := T*( 1 - ( math.Acos((1-g)/e) - math.Sqrt( g*(2 - b*o - g) ) )/math.Pi )

	j.Time += t

	return flightTimeMsg(j, t)
}

func ButlerFlightTime(j *J) string {
	v,thetadash := j.Velocity,j.ThetaDash
	if v == 0 {
		return
	}
	v0 := v*math.Cos(thetadash)
	g := Mu/(R*R)
	h := R*v0*v0/(2*R*g - v0*v0)
	a := v0*v0*R
	b := v0*v0 - (2*R*g)
	u := a + b*h
	w := R + h
	l := a - b*R
	p := (2*b*h + a + b*R)/l

	hmax := math.Copysign(2,w)*(math.Sqrt(u*w)/b + math.Asin(p)*l/(2*b*math.Sqrt(-b)))
	h0 := 2*(math.Sqrt(a*R)/b + math.Asin((a+b*R)/l)*l/(2*b*math.Sqrt(-b)))

	t := hmax - h0

	j.Time += t

	return flightTimeMsg(j, t)
}

func VondrakFlightTime(j *J) string {
	v,thetadash := j.Velocity,j.ThetaDash
	if v == 0 {
		return
	}
	vr := v*math.Cos(thetadash)
	g := Mu/(R*R)
	z := vr*vr/(g*R)

	t := 2*vr*(1 + (math.Pi/2 + math.Asin(z - 1))/math.Sqrt(z*(2-z)))/(g*(2-z))

	j.Time += t

	return flightTimeMsg(j, t)
}

func bruteforce(j *J) string {
	v,thetadash := j.Velocity,j.ThetaDash
	a := R/(2 - v*v*R/Mu)
	s := math.Sin(thetadash)*math.Sin(thetadash)*R*(2*a - R)/(a*a)
	e := math.Sqrt( 1 - s )
	o := math.Acos((a*s/R - 1)/e)
	E := 2*math.Atan(math.Sqrt((1-e)/(1+e))*math.Tan(o/2))
	M := E - e*math.Sin(E)
	T := 2*math.Pi*math.Sqrt(a*a*a/Mu)

	M = 2*math.Pi - 2*M
	t := M*T/(2*math.Pi)

	j.Time += t

	return flightTimeMsg(j, t)
}

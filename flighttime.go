package bouncing

import "math"

func FlightTime(j *J) {
	v,thetadash := j.V,j.ThetaDash
	if v == 0 {
		j.T = 0
		return
	}
	b := v*v*R/MU
	g := 2 - b
	o := math.Sin(thetadash)*math.Sin(thetadash)
	e := math.Sqrt(1-b*g*o)
	T := 2*math.Pi*math.Sqrt(R*R*R/(MU*g*g*g))

	t := T*( 1 - ( math.Acos((1-g)/e) - math.Sqrt( g*(2 - b*o - g) ) )/math.Pi )

	j.T = t
}

func ButlerFlightTime(j *J) {
	v,thetadash := j.V,j.ThetaDash
	if v == 0 {
		j.T = 0
		return
	}
	v0 := v*math.Cos(thetadash)
	g := MU/(R*R)
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

	j.T = t
}

func VondrakFlightTime(j *J) {
	v,thetadash := j.V,j.ThetaDash
	if v == 0 {
		j.T = 0
		return
	}
	vr := v*math.Cos(thetadash)
	g := MU/(R*R)
	z := vr*vr/(g*R)

	t := 2*vr*(1 + (math.Pi/2 + math.Asin(z - 1))/math.Sqrt(z*(2-z)))/(g*(2-z))

	j.T = t
}

func bruteforce(j *J) {
	v,thetadash := j.V,j.ThetaDash
	a := R/(2 - v*v*R/MU)
	s := math.Sin(thetadash)*math.Sin(thetadash)*R*(2*a - R)/(a*a)
	e := math.Sqrt( 1 - s )
	o := math.Acos((a*s/R - 1)/e)
	E := 2*math.Atan(math.Sqrt((1-e)/(1+e))*math.Tan(o/2))
	M := E - e*math.Sin(E)
	T := 2*math.Pi*math.Sqrt(a*a*a/MU)

	M = 2*math.Pi - 2*M
	t := M*T/(2*math.Pi)

	j.T = t
}

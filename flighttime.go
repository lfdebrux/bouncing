package bouncing

import "math"

type F func(v,thetadash float64) (t float64)

func FlightTime(v,thetadash float64) (t float64) {
	b := v*v*R/MU
	g := 2 - b
	o := math.Sin(thetadash)*math.Sin(thetadash)
	c := b*o
	s := b*g*o
	e := math.Sqrt(1-s)
	T := 2*math.Pi*math.Sqrt(R*R*R/(MU*g*g*g))

	t = T*( 1 - ( math.Acos((1-g)/e) - math.Sqrt( s*(e*e-(c-1)*(c-1)) )/c )/math.Pi )
	return t
}

func ButlerFlightTime(v,thetadash float64) (t float64) {
	v0 := v*math.Cos(thetadash)
	h := R*v0*v0/(2*MU/R - v0*v0)
	a := v0*v0*R
	b := v0*v0 - (2*MU/R)
	u := a + b*h
	w := R + h
	l := a - b*R
	p := (2*b*h + a + b*R)/l

	hmax := math.Copysign(2,w)*(math.Sqrt(u*w)/b + math.Asin(p)*l/(2*b*math.Sqrt(-b)))
	h0 := 2*(math.Sqrt(a*R)/b + math.Asin((a+b*R)/l)*l/(2*b*math.Sqrt(-b)))

	t = hmax - h0

	return t
}

func VondrakFlightTime(v,thetadash float64) (t float64) {
	vr := v*math.Cos(thetadash)
	g := MU/(R*R)
	z := vr*vr/(g*R)

	t = 2*vr*(1 + (math.Pi/2 + math.Asin(z - 1))/math.Sqrt(z*(2-z)))/(g*(2-z))

	return t
}
func bruteforce(v,thetadash float64) (t float64) {
	a := R/(2 - v*v*R/MU)
	s := math.Sin(thetadash)*math.Sin(thetadash)*R*(2*a - R)/(a*a)
	e := math.Sqrt( 1 - s )
	o := math.Acos((a*s/R - 1)/e)
	E := 2*math.Atan(math.Sqrt((1-e)/(1+e))*math.Tan(o/2))
	M := E - e*math.Sin(E)
	T := 2*math.Pi*math.Sqrt(a*a*a/MU)

	M = 2*math.Pi - 2*M
	t = M*T/(2*math.Pi)

	return t
}

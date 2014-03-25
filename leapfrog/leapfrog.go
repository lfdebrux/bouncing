package leapfrog

import "code.google.com/p/gomat/vec"
import "math"

import . "github.com/lfdebrux/bouncing"

const dt = 5e-4

type Point struct {
	r vec.DenseVector
	v vec.DenseVector
}

func NewFromJump(j *J) *Point {
	r,v := vec.New(3),vec.New(3)

	cosbeta := math.Cos(j.Beta)
	sinbeta := math.Sin(j.Beta)
	cosphi := math.Cos(j.Phi)
	sinphi := math.Sin(j.Phi)

	r[0] = R*sinphi*sinbeta
	r[1] = R*cosphi
	r[2] = R*sinphi*cosbeta

	rdot := j.V*math.Cos(j.ThetaDash) // v along normal (direction of increasing r)
	Rbetadot := -j.V*math.Sin(j.ThetaDash)*math.Sin(j.Psi) // R*angular v along direction of increasing beta
	Rphidot := -j.V*math.Sin(j.ThetaDash)*math.Cos(j.Psi) // R*angular v along direction of increasing phi

	v[0] = rdot*sinbeta*sinphi + Rbetadot*cosbeta*sinphi + Rphidot*sinbeta*cosphi
	v[1] = rdot*cosphi - Rphidot*sinphi
	v[2] = rdot*cosbeta*sinphi - Rbetadot*sinbeta*sinphi + Rphidot*cosbeta*cosphi

	return &Point{r,v}
}

func (p *Point) Beta() float64 {
	return math.Atan2(p.r[0],p.r[2])
}

func (p *Point) Phi() float64 {
	return math.Acos(p.r[1]/math.Sqrt(p.r.Dot(p.r)))
}

func (p *Point) LeapFrog() {
	for i := range p.r {
		p.r[i] += p.v[i]*dt
	}

	r2 := p.r.Dot(p.r)
	a := -dt*Mu/(r2*math.Sqrt(r2))

	for i := range p.v {
		p.v[i] += p.r[i]*a
	}
}

func (p *Point) LeapFrogUntil() (n float64) {
	for p.r.Dot(p.r) >= R*R {
		p.LeapFrog()
		n++
	}
	return n
}

func (p *Point) MaxAltitude() (h float64) {
	r2 := p.r.Dot(p.r)
	for r2 >= R*R {
		p.LeapFrog()
		r2 = p.r.Dot(p.r)
		if p.r.Dot(p.r) > h {
			h = r2
		}
	}
	return h
}
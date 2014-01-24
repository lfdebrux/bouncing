package leapfrog

import "testing"

import "math"
import "math/rand"

import . "github.com/lfdebrux/bouncing"

const NUM = 10
const TOL = 1e-2

func diff(a,b float64) float64 {
	return math.Abs(a-b)
}

func almosteq(b,a float64) bool {
	scale := math.Abs(a)
	if scale == 0 {
		scale = TOL
	}
	if diff(a,b) < TOL*scale {
		return true
	}
	return false
}

func simplecase(v0,thetadash float64) (p *Point) {
	return NewFromJump(&J{P:&P{Beta:0,Phi:math.Pi/2},V:v0,Psi:0,ThetaDash:thetadash})
}

func TestBetaPhi(t *testing.T) {
	p := simplecase(300,math.Pi/4)

	if !almosteq(p.Beta(),0.0) {
		t.Errorf("expected p.Beta = %f, instead got %f",0.0,p.Beta())
	}
	if !almosteq(p.Phi(),math.Pi/2) {
		t.Errorf("expected p.Phi = %f, instead got %f",math.Pi/2,p.Phi())
	}
}

func TestAgainstKepler(t *testing.T) {
	for i := 0; i < NUM; i++ {
		v0,thetadash := rand.Float64()*700.0,rand.Float64()*math.Pi/2
		p := simplecase(v0,thetadash)

		tleap := p.LeapFrogUntil()*DT
		tkep := FlightTime(v0,thetadash)

		if math.Abs(tleap-tkep) > TOL/DT {
			t.Log(TOL/DT)
			t.Log(tleap - tkep)
			t.Fatalf("expected n~=%f, instead got %f",tleap,tkep)
		}
	}
}

func TestAgainstPositionJump(t *testing.T) {
	NUM := 10
	FAIL := 0
	for i := 0; i < NUM; i++ {
		v0,thetadash := rand.Float64()*700.0,rand.Float64()*math.Pi/2
		p := simplecase(v0,thetadash)

		t.Logf("initial conditions %d v0 = %f thetadash = %f",i,v0,thetadash)
		t.Logf("initial position %d (phi,beta) = (%f,%f)",i,p.Phi(),p.Beta())

		p.LeapFrogUntil()
		phileap,betaleap := p.Phi(),p.Beta()
		phipos,betapos := VondrakPositionJump(math.Pi/2,0,v0,0,thetadash)

		t.Logf("final position %d (phi,beta): leapfrog (%f,%f); jump (%f,%f)",i,phileap,betaleap,phipos,betapos)

		if !almosteq(betapos,betaleap) || !almosteq(phipos,phileap) {
			t.Logf("jump %d not within required accuracy",i)
			FAIL++
			t.Fail()
		}
	}
	t.Logf("%d jumps not within required accuracy",FAIL)
}

func TestNewFromJump(t *testing.T) {
	v0,thetadash := 300.0,1.3
	p := simplecase(v0,thetadash)

	if expect := R; !almosteq(p.r[2],expect) {
		t.Log(p.r[2] - expect)
		t.Errorf("expected p.r[2] = %f, instead got p.r=%v p.v=%v",expect,p.r,p.v)
	}
	if expect := v0*math.Sin(thetadash); !almosteq(p.v[1],expect) {
		t.Log(p.v[1] - expect)
		t.Errorf("expected p.v[1] = %f, instead got p.r=%v p.v=%v",expect,p.r,p.v)
	}
	if expect := v0*math.Cos(thetadash); !almosteq(p.v[2],expect) {
		t.Log(p.v[2] - expect)
		t.Errorf("expected p.v[2] = %f, instead got p.r=%v p.v=%v",expect,p.r,p.v)
	}
}

func TestLeapFrogUntil(t *testing.T) {
	p := simplecase(300,math.Pi)
	if n := p.LeapFrogUntil(); n > 1 {
		t.Errorf("expected LeapFrogUntil to end immediately, but it took %v steps",n)
	}
}

func BenchmarkLeapFrog(b *testing.B) {
	p := simplecase(300,math.Pi/4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.LeapFrog()
	}
}

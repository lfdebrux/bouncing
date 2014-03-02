package bouncing

import "testing"

import (
	"math"
	"math/rand"
	"time"
	"github.com/lfdebrux/bouncing/hist"
)

func maxwellian(a,x float64) float64 {
	return math.Sqrt(2/math.Pi)*(a*a*a)*(x*x)*math.Exp(-((a*x)*(a*x))/2)
}

func maxwellianCDF(a,x float64) float64 {
	return math.Erf(a*x/math.Sqrt2) - math.Sqrt(2/math.Pi)*a*x*math.Exp(-((a*x)*(a*x))/2)
}

func TestMaxwellianInverse(t *testing.T) {
	InitMaxwellian()
	defer FreeMaxwellian()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < NUM; i++ {
		a := rand.Float64()
		v := 10*rand.Float64()

		p := maxwellianCDF(a,v)
		av_test := Q.Eval(p)

		if !almosteq(av_test, a*v) {
			t.Fatalf("Q(maxwellianCDF(%v,%v)) should be %v, instead got %v, p = %v",a,v,a*v,av_test,p)
		}
	}
}

func averageAccumulator() (func(x float64) (float64,float64)) {
	var k,Qk,mean,stddev float64

	return func(x float64) (float64,float64) {
		if math.IsNaN(x) {
			return mean,stddev
		}
		if k == 0 {
			k++
			mean = x
		} else {
			k++
			Qk += (x-mean)*(x-mean)*(k-1)/k
			mean += (x-mean)/k
			stddev += math.Sqrt(Qk/(k-1))
		}

		return mean,stddev
	}
}

func TestMaxwellianMode(t *testing.T) {
	InitMaxwellian()
	defer FreeMaxwellian()

	rand.Seed(time.Now().UnixNano())

	b := hist.NewBins(0,8,0.1)

	for i := 0; i < NUM; i++ {
		v := Q.Eval(rand.Float64())
		b.Bin(v)
	}

	var max,mode float64
	for i,n := range b.Binned {
		if n > max {
			max = n
			mode = b.Binmap[i]
		}
	}

	if !almosteq(mode,1.4/*math.Sqrt2*/) {
		t.Errorf("Mode of Maxwell-Boltzmann distribution should be sqrt(2)=%f, instead got %f (%v counts)",math.Sqrt2,mode,max)
	}
}

func TestMaxwellianRMS(t *testing.T) {
	InitMaxwellian()
	defer FreeMaxwellian()

	rand.Seed(time.Now().UnixNano())

	avg := averageAccumulator()

	for i := 0; i < NUM; i++ {
		v := Q.Eval(rand.Float64())
		avg(v*v)
	}

	mean,sd := avg(math.NaN())

	t.Log(mean,sd)

	if !almosteq(mean,3) {
		t.Errorf("Mean Square of Maxwell-Boltzmann distribution should be 3, instead got %f",mean)
	}
}

func TestMaxwellianMean(t *testing.T) {
	InitMaxwellian()
	defer FreeMaxwellian()

	rand.Seed(time.Now().UnixNano())

	avg := averageAccumulator()

	for i := 0; i < NUM; i++ {
		v := Q.Eval(rand.Float64())
		avg(v)
	}

	mean,sd := avg(math.NaN())

	t.Log(mean,sd)

	if !almosteq(mean, math.Sqrt(8/math.Pi)) {
		t.Errorf("Mean of Maxwell-Boltzmann distribution should be sqrt(8/Pi)=%f, instead got %f",math.Sqrt(8/math.Pi),mean)
	}
}

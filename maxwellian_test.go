package bouncing

import "testing"

import (
	"math"
	"math/rand"
	"time"
	"github.com/lfdebrux/bouncing/hist"
)

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

func TestUnInitedMaxwellian(t *testing.T) {
	if Q != nil {
		FreeMaxwellian()
	}
	defer FreeMaxwellian()

	RandVelocity(1,0)
}

func BenchmarkRandVelocity(b *testing.B) {
	InitMaxwellian()
	defer FreeMaxwellian()

	m := Mass[Water]
	lat := 0.0

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RandVelocity(m,lat)
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

	t.Log(mode,max)

	if !almosteq(mode,1.4/*math.Sqrt2*/) {
		t.Errorf("Mode of Maxwell-Boltzmann distribution should be sqrt(2)=%f, instead got %f",math.Sqrt2,mode)
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

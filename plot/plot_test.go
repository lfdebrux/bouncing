package plot

import "testing"
import "math"

func rangetesthelper(t *testing.T, start,stop,step float64, expected []float64) {
	j := 0
	for _,i := range Range(start,stop,step) {
		t.Log(j,i)
		if math.Abs(i - expected[j]) > 1e-13 {
			t.Errorf("%dth element of Range(%f,%f,%f) is %f, expected %f",j,start,stop,step,i,expected[j])
		}
		j++
	}
	if j != len(expected) {
		t.Errorf("expected %d elements from Range(%f,%f,%f), but it stopped at %d",len(expected),start,stop,step,j-1)
	}
}

func TestRange(t *testing.T) {
	test := []float64{0,1,2,3,4,5,6,7,8,9,10}
	rangetesthelper(t,0,10,1,test)
}

func TestRangeWithDoubleStep(t *testing.T) {
	test := []float64{0,2,4,6,8,10,12,14,16,18,20}
	rangetesthelper(t,0,20,2,test)
}

func TestRangeNonZeroStart(t *testing.T) {
	test := []float64{10,11,12,13,14,15,16,17,18,19,20}
	rangetesthelper(t,10,20,1,test)
}

func TestRangeNonZeroStartWithDoubleStep(t *testing.T) {
	test := []float64{10,12,14,16,18,20,22,24,26,28,30}
	rangetesthelper(t,10,30,2,test)
}

func TestRangePiSteps(t *testing.T) {
	test := []float64{0,math.Pi/4,math.Pi/2,3*math.Pi/4,math.Pi}
	rangetesthelper(t,0,math.Pi,math.Pi/4,test)
}

func TestRangePiSmallSteps(t *testing.T) {
	test := make([]float64,100)
	for i := 0; i < 100; i++ {
		test[i] = math.Pi*float64(i)/(100-1)
	}
	rangetesthelper(t,0,math.Pi,math.Pi/(100-1),test)
}

func TestRangeVerySmallSteps(t *testing.T) {
	test := []float64{0,0.001,0.002,0.003,0.004,0.005,0.006,0.007,0.008,0.009,0.010}
	rangetesthelper(t,0,0.01,0.001,test)	
}
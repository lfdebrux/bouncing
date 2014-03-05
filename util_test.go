package bouncing

import "math"

const NUM = 1e6
const TOL = 1e-2

func newJ() *J {
	return &J{P:new(P)}
}

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

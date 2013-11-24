package bouncing

import "math"

const NUM = 1000000
const TOL = 1e-6

func diff(a,b float64) float64 {
	return math.Abs(a-b)
}

func almosteq(a,b,tol float64) bool {
	if diff(a,b) < tol {
		return true
	}
	return false
}

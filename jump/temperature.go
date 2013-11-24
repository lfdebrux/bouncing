package jump

import "math"

const T0 = 151.0 // K
const T1 = 161.7 // K
const Tn = 0.59

func Temperature(phi float64) float64 {
	return T0 + T1*math.Pow(math.Cos(phi-math.Pi/2),Tn)
}

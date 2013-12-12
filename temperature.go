package bouncing

import "math"

type TemperatureFunc func(float64) float64

func ButlerTemperature(phi float64) float64 {
	return 151 + 161.7*math.Pow(math.Cos(phi-math.Pi/2),0.59)
}

func VondrakTemperature(solarzenith float64) float64 {
	return 280 * math.Pow(math.Cos(solarzenith),0.25) + 100
}

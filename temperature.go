package bouncing

import "math"

const (
	ButlerT0 = 151.0 // Kelvin
	ButlerT1 = 161.7 // Kelvin
)

func ButlerTemperature(j *J) *Lost {
	j.Temperature = ButlerT0 + ButlerT1*math.Pow(math.Cos(j.Phi-math.Pi/2),0.59)
	return nil
}

func VondrakTemperature(j *J) *Lost {
	j.Temperature = 280 * math.Pow(math.Cos(j.SolarZenith),0.25) + 100
	return nil
}

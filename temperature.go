package bouncing

import (
	_ "fmt"
	"math"
	_ "math/rand"
)

const (
	ButlerT0 = 151.0 // Kelvin
	ButlerT1 = 161.7 // Kelvin
)

func ButlerTemperature(j *J) {
	j.Temperature = ButlerT0 + ButlerT1*math.Pow(math.Cos(j.Phi-math.Pi/2),0.59)
}

func VondrakTemperature(j *J) string {
	j.Temperature = 280 * math.Pow(math.Cos(j.SolarZenith),0.25) + 100
	// Temperature is complex if SolarZenith > pi/2, so Pow retuns NaN
	if j.SolarZenith > math.Pi/2 {
		return "e NaN VondrakTemperature, SolarZenith is greater than pi/2"
	}
	return ""
}

// func TemperatureResidence(j *J) string {
// 	// generate a random time to thermally accomodate before being desorped
// 	decayrate := // TODO
// 	t = -math.Log(1 - rand.Float64())/decayrate
// 	j.Time += t
// 	return fmt.Sprintln("r", j.Time, j.Phi, j.Beta, t, "thermally accomodating")
// }
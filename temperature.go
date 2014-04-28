package bouncing

import (
	"fmt"
	"math"
	"math/rand"
)

const timePerDeg = LunarDay/360 // time to turn 1 degrees

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
		j.Type = Error
		return "e NaN VondrakTemperature, SolarZenith is greater than pi/2\n"
	}
	return ""
}

func KillenTemperature(j *J) {
	if j.SolarZenith < math.Pi/2 {
		j.Temperature = 280 * math.Pow(math.Cos(j.SolarZenith),0.25) + 100
	} else if j.SolarZenith > math.Pi/2 {
		j.Temperature = 100
	}
}

func ThermalResidence(j *J) string {
	// generate a random time to thermally accomodate before being desorped
	rate := 1e13*math.Exp(-5802.26/j.Temperature)
	t := -math.Log(1 - rand.Float64())/rate
	if t > timePerDeg { // time to turn 1 degrees
		t = timePerDeg // clamp because rotation (temperature) affects desorption rate
	}
	j.Time += t
	j.Beta += 2*math.Pi*t/LunarDay // rotation
	return fmt.Sprintln("r", j.Time, j.Phi, j.Beta, t, "thermally accomodating")
}

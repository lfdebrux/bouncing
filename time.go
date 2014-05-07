package bouncing

import "fmt"
import "math"

const timeperrad = 375704 // LunarDay/2pi // seconds/rad

func SolarHourAngle(time, phi, beta float64) float64 {
	hr := beta + 2*math.Pi*time/LunarDay
	hr = math.Mod(hr,2*math.Pi) - math.Pi
	return hr
}

func ZenithFromTime(time, phi, beta float64) (solarzenith float64) {
	// Calculate the solar zenith angle for a given time and position.

	// currently a very naive formula, assuming that the lunar equator
	// is aligned with the Sun-Moon plane, that the lunar rotation is constant,
	// and that it is exactly midnight at colatitude Pi/2, longitude 0, time 0.

	hr := SolarHourAngle(time, phi, beta)
	return math.Acos( math.Cos(phi - math.Pi/2)*math.Cos(hr) )
}

func TimeToSunrise(time, phi, beta float64) float64 {
	hr := SolarHourAngle(time,phi,beta)/(2*math.Pi) + 0.25
	if hr > 0.50 {
		hr -= 1.0
	}
	return -LunarDay*hr
}

func VondrakZenithOnly(j *J) {
	j.SolarZenith = math.Acos( math.Sin(j.Phi)*math.Cos(j.Beta) )
}

// Zenith angle with Sun fixed above longitude 0
// skip ahead if on nightside
func VondrakZenith(j *J) string {
	if math.Pi/2 < j.Beta && j.Beta < 3*math.Pi/2 {
		j.Time += (3*math.Pi/2 - j.Beta)*timeperrad
		j.Beta = 3*math.Pi/2 + 1e-5 // ensure > 3pi/2
		j.Type = NightSide
		j.SolarZenith = math.Acos( math.Sin(j.Phi)*math.Cos(j.Beta) )
		return fmt.Sprintln("d", j.Time, j.Phi, j.Beta)
	}
	j.SolarZenith = math.Acos( math.Sin(j.Phi)*math.Cos(j.Beta) )
	return ""
}

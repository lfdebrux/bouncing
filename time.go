package bouncing

import "math"

const LUNARDAY = 2.3606e6 // seconds
const TIMEPERRAD = 375704 // LUNARDAY/2pi // seconds/rad

func SolarHourAngle(time, phi, beta float64) float64 {
	hr := beta + 2*math.Pi*time/LUNARDAY
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
	return -LUNARDAY*hr
}

// Zenith angle with Sun fixed above longitude 0
// skip ahead if on nightside
func VondrakZenith(j *J) *Lost {
	j.SolarZenith = math.Acos( math.Sin(j.Phi)*math.Cos(j.Beta) )
	if j.SolarZenith > math.Pi/2 {
		j.Time += (j.Beta - math.Pi/2)*TIMEPERRAD
		j.Beta = 3*math.Pi/2
		j.SolarZenith = math.Pi/2
	}
	return nil
}

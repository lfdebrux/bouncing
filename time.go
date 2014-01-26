package bouncing

import "math"

const LUNARDAY = 2.3606e6 // seconds

func SolarHourAngle(time, phi, beta float64) float64 {
	hr := beta + 2*math.Pi*time/LUNARDAY
	hr = math.Mod(hr,2*math.Pi) - math.Pi
	return hr
}

func ZenithFromTime(time, phi, beta float64) (solarzenith float64) {
	// Calculate the solar zenith angle for a given time and position.

	// currently a very naive formula, assuming that the lunar equator
	// is aligned with the orbital plane, that the lunar rotation is constant,
	// and that it is exactly midnight at colatitude Pi/2, longitude 0, time 0.

	hr := SolarHourAngle(time, phi, beta)
	return math.Acos( math.Cos(phi - math.Pi/2)*math.Cos(hr) )
}

func TimeToSunrise(time, phi, beta float64) float64 {
	hr := SolarHourAngle(time,phi,beta)/(2*math.Pi) + 0.25
	if hr > 0.50 {
		hr -= 1.0
	}
	return LUNARDAY*hr
}

package bouncing

import "math"

const LUNARDAY = 2.3606E6 // seconds

func zenithFromTime(time, phi, beta float64) (solarzenith float64) {
	// Calculate the solar zenith angle for a given time and position.

	// currently a very naive formula, assuming that the lunar equator
	// is aligned with the orbital plane, that the lunar rotation is constant,
	// and that the sun is directly overhead latitude 0, longitude 0, time 0.

	hr := beta + 2*math.Pi*time/LUNARDAY // solar hour angle
	return math.Acos( math.Cos(phi - math.Pi/2)*math.Cos(hr) )
}

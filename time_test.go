package bouncing

import "testing"

func BenchmarkVondrakZenith(b *testing.B) {
	j := newJ()
	for i := 0; i < b.N; i++ {
		VondrakZenith(j)
	}
}

func BenchmarkVondrakSunrise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VondrakSunrise(&J{P:&P{SolarZenith:3}})
	}
}

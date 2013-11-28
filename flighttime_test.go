package bouncing

import "testing"

func BenchmarkFlightTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FlightTime(300,1.74)
	}
}

func BenchmarkButlerFlightTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ButlerFlightTime(300,1.74)
	}
}

func BenchmarkVondrakFlightTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VondrakFlightTime(300,1.74)
	}
}

func TestButlerFlightTimeZero(t *testing.T) {
	if time := ButlerFlightTime(0,0); time != 0 {
		t.Errorf("Expected flight time with v=0 using Butler equation to be 0, instead got %f",time)
	}
}

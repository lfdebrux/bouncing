package jump

import "testing"

func TestButlerFlightTimeZero(t *testing.T) {
	if time := ButlerFlightTime(0,0); time != 0 {
		t.Errorf("Expected flight time with v=0 using Butler equation to be 0, instead got %f",time)
	}
}

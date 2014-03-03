package bouncing

import "testing"

func TestJump(t *testing.T) {
	InitMaxwellian()
	defer FreeMaxwellian()

	p := RandParticle(Water)
	Jump(p)
}

func TestNewJump(t *testing.T) {
	InitMaxwellian()
	defer FreeMaxwellian()

	jump := NewJump(ButlerTemperature,RandVelocity,RandDirection,ButlerPositionJump,FlightTime)
	p := RandParticle(Water)

	jump(p)
}

func BenchmarkJump(b *testing.B) {
	InitMaxwellian()
	defer FreeMaxwellian()
	
	p := RandParticle(Water)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Jump(p)
	}
}

func BenchmarkNewJump(b *testing.B) {
	InitMaxwellian()
	defer FreeMaxwellian()
	
	jump := NewJump(ButlerTemperature,RandVelocity,RandDirection,ButlerPositionJump,FlightTime)

	p := RandParticle(Water)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jump(p)
	}
}

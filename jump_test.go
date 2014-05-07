package bouncing

import "testing"

func TestJump(t *testing.T) {
	InitMaxwellian()
	defer FreeMaxwellian()

	p := new(P)
	
	Jump(p)
}

func TestVondrakJump(t *testing.T) {
	InitMaxwellian()
	defer FreeMaxwellian()

	p := new(P)
	
	VondrakJump(p)
}

func TestNewJumpSimple(t *testing.T) {
	InitMaxwellian()
	defer FreeMaxwellian()

	jump := NewJumpSimple(ButlerTemperature,RandVelocity,RandDirection,ButlerPositionJump,FlightTime)
	p := new(P)
	

	jump(p)
}

func BenchmarkJump(b *testing.B) {
	InitMaxwellian()
	defer FreeMaxwellian()
	
	p := new(P)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Jump(p)
	}
}

func BenchmarkVondrakJump(b *testing.B) {
	InitMaxwellian()
	defer FreeMaxwellian()
	
	p := new(P)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		VondrakJump(p)
	}
}

func BenchmarkNewJumpSimple(b *testing.B) {
	InitMaxwellian()
	defer FreeMaxwellian()
	
	jump := NewJumpSimple(ButlerTemperature,RandVelocity,RandDirection,ButlerPositionJump,FlightTime)

	p := new(P)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jump(p)
	}
}

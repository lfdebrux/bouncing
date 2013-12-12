package bouncing

import "testing"

func BenchmarkJump(b *testing.B) {
	InitMaxwellian()
	defer FreeMaxwellian()
	
	p := RandParticle(Water)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Jump(p)
	}
}
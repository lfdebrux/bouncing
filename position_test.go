package bouncing

import "testing"

func BenchmarkPositionJump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ButlerPositionJump(1.35,3.1,463,2.98,1.7)
	}
}

func BenchmarkVondrakPositionJump(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VondrakPositionJump(1.35,3.1,463,2.98,1.7)
	}
}

// func TestPositionWithLeapfrog(t *testing.T) {
	// 
// }

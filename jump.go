package bouncing

type J struct {
	*P
	Velocity,Psi,ThetaDash float64
	Temperature float64
	FlightTime float64
}

type JumpMethod func(*J) *Lost
type JumpFunc func(*P) (*J,*Lost)

func Jump(p *P) (*J,*Lost) {
	j := &J{P:p}
	ButlerTemperature(j)
	RandVelocity(j)
	RandDirection(j)
	ButlerPositionJump(j)
	FlightTime(j)

	j.Time += j.FlightTime

	return j,IsLost(j)
}

func NewJump(funcs ...JumpMethod) JumpFunc {
	return func(p *P) (*J,*Lost) {
		j := &J{P:p}
		
		for _,f := range funcs {
			f(j)
		}

		j.Time += j.FlightTime

		return j,IsLost(j)
	}
}

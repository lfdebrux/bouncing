package bouncing

type J struct {
	*P
	V,Psi,ThetaDash float64
	Temperature float64
	T float64
}

type JumpMethod func(*J)
type JumpFunc func(*P) (*J,*Lost)

func Jump(p *P) (*J,*Lost) {
	j := &J{P:p}
	ButlerTemperature(j)
	RandVelocity(j)
	RandDirection(j)
	ButlerPositionJump(j)
	FlightTime(j)

	j.Time += j.T

	if lost := j.IsLost(); lost != nil {
		p = RandParticle(p.Type)
		return nil,lost
	}

	return j,nil
}

func NewJump(funcs ...JumpMethod) JumpFunc {
	return func(p *P) (*J,*Lost) {
		j := &J{P:p}
		
		for _,f := range funcs {
			f(j)
		}

		j.Time += j.T

		if lost := j.IsLost(); lost != nil {
			p = RandParticle(p.Type)
			return nil,lost
		}

		return j,nil
	}
}

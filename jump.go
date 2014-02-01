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
	j.Psi,j.ThetaDash = RandDirection()
	j.Phi,j.Beta = ButlerPositionJump(j.Phi,j.Beta,j.V,j.Psi,j.ThetaDash)
	FlightTime(j)

	j.Time += j.T

	if lost := j.IsLost(); lost != nil {
		p = RandParticle(p.Type)
		return nil,lost
	}

	return j,nil
}

func NewJump(tm JumpMethod,rd RandDirectionFunc,pj PositionJumpFunc,ft JumpMethod) JumpFunc {
	return func(p *P) (*J,*Lost) {
		j := &J{P:p}
		tm(j)
		RandVelocity(j)
		j.Psi,j.ThetaDash = rd()
		j.Phi,j.Beta = pj(j.Phi,j.Beta,j.V,j.Psi,j.ThetaDash)
		ft(j)

		j.Time += j.T

		if lost := j.IsLost(); lost != nil {
			p = RandParticle(p.Type)
			return nil,lost
		}

		return j,nil
	}
}

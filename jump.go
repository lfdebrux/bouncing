package bouncing

type J struct {
	*P
	V,Psi,ThetaDash float64
	Temperature float64
	T float64
}

type JumpFunc func(*P) (*J,*Lost)

func Jump(p *P) (*J,*Lost) {
	j := &J{P:p}
	j.Temperature = ButlerTemperature(j.Phi,j.Beta)
	j.V = RandVelocity(Mass[j.Type],j.Temperature)
	j.Psi,j.ThetaDash = RandDirection()
	j.Phi,j.Beta = ButlerPositionJump(j.Phi,j.Beta,j.V,j.Psi,j.ThetaDash)
	j.T = FlightTime(j.V,j.ThetaDash)

	j.Time += j.T

	if lost := j.IsLost(); lost != nil {
		p = RandParticle(p.Type)
		return nil,lost
	}

	return j,nil
}

func NewJump(tm TemperatureFunc,rd RandDirectionFunc,pj PositionJumpFunc,ft FlightTimeFunc) JumpFunc {
	return func(p *P) (*J,*Lost) {
		j := &J{P:p}
		j.Temperature = tm(j.Phi,j.Beta)
		j.V = RandVelocity(Mass[j.Type],j.Temperature)
		j.Psi,j.ThetaDash = rd()
		j.Phi,j.Beta = pj(j.Phi,j.Beta,j.V,j.Psi,j.ThetaDash)
		j.T = ft(j.V,j.ThetaDash)

		if lost := j.IsLost(); lost != nil {
			p = RandParticle(p.Type)
			return nil,lost
		}

		return j,nil
	}
}

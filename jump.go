package bouncing

type J struct {
	*P
	V,Psi,ThetaDash float64
	T float64
}

func Jump(p *P) (*J,*Lost) {
	j := &J{P:p}
	j.V = RandVelocity(Mass[j.Type],j.Phi)
	j.Psi,j.ThetaDash = RandDirection()
	j.Phi,j.Beta = PositionJump(j.Phi,j.Beta,j.V,j.Psi,j.ThetaDash)
	j.T = FlightTime(j.V,j.ThetaDash)

	if lost := j.IsLost(); lost != nil {
		p = RandParticle(p.Type)
		return nil,lost
	}

	return j,nil
}

func NewJump(rd RandDirectionFunc,pj PositionJumpFunc,ft FlightTimeFunc) func(*P) {
	return func(p *P) (*J,*Lost) {
		j := &J{P:p}
		j.V = RandVelocity(Mass[j.Type],j.Phi)
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

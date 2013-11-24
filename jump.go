package bouncing

type J struct {
	*P
	V,Psi,ThetaDash float64
	T float64
}

func Jump(p *P) (*J,*Lost) {
	j := &J{P:p}
	j.V = RandVelocity(Mass[j.Type],j.Phi)
	j.Psi,j.ThetaDash = ButlerRandDirection()
	j.Phi,j.Beta = PositionJump(j.Phi,j.Beta,j.V,j.Psi,j.ThetaDash)
	j.T = ButlerFlightTime(j.V,j.ThetaDash)

	if lost := j.IsLost(); lost != nil {
		p = RandParticle(p.Type)
		return nil,lost
	}

	return j,nil
}

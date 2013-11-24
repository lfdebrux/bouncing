package jump

import "github.com/lfdebrux/bouncing/particle"

type J struct {
	*particle.Particle
	V,Psi,ThetaDash float64
	T float64
}

func Jump(p *particle.Particle) (*J,*Lost) {
	j := &J{Particle:p}
	j.V = RandVelocity(particle.Mass[j.Type],j.Phi)
	j.Psi,j.ThetaDash = ButlerRandDirection()
	j.Phi,j.Beta = PositionJump(j.Phi,j.Beta,j.V,j.Psi,j.ThetaDash)
	j.T = ButlerFlightTime(j.V,j.ThetaDash)

	if lost := j.IsLost(); lost != nil {
		p = particle.RandParticle(p.Type)
		return nil,lost
	}

	return j,nil
}

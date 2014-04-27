package bouncing

type JumpType int

const (
	_ = iota

	isJump JumpType = iota	// marker
	NightSide

	isLost					// marker
	ThermalEscape
	Photodestruction
	Capture
	Error
)

type J struct {
	*P
	Type JumpType
	Velocity,Psi,ThetaDash float64
	Temperature float64
	FlightTime float64
}

type JumpMethodSimple func(*J)
type JumpMethod func(*J) string
type JumpFunc func(*P) (*J, string)

func lift(f JumpMethodSimple) JumpMethod {
	return func(j *J) string {
		f(j)
		return ""
	}
}

func liftAll(fs ...interface{}) []JumpMethod {
	out := make([]JumpMethod, len(fs))
	for i, f := range fs {
		switch f := f.(type) {
		case func(*J):
			out[i] = lift(f)
		case func(*J) string:
			out[i] = f
		default:
			panic("arguments must be JumpMethod or JumpMethodSimple")
		}
	}
	return out
}

func Jump(p *P) (*J, string) {
	j := &J{P:p}
	ButlerTemperature(j)
	RandVelocity(j)
	RandDirection(j)
	ButlerPositionJump(j)
	FlightTime(j)

	msg := CheckLost(j)
	if msg == "" {
		msg = CaptureButler(j)
	}

	return j, msg
}

func JumpWithVondrak(p *P) (*J,  string) {
	j := &J{P:p}
	VondrakZenith(j)
	VondrakTemperature(j)
	RandVelocity(j)
	RandDirection(j)
	ButlerPositionJump(j)
	FlightTime(j)

	msg := CheckLost(j)
	if msg == "" {
		msg = CaptureVondrak(j)
	}

	return j, msg
}

func ButlerJump(p *P) (*J,  string) {
	j := &J{P:p}
	ButlerTemperature(j)
	RandVelocity(j)
	ButlerRandDirection(j)
	ButlerPositionJump(j)
	VondrakFlightTime(j) // ButlerFlightTime is NaNy

	msg := CheckLost(j)
	if msg == "" {
		msg = CaptureButler(j)
	}

	return j, msg
}

func VondrakJump(p *P) (*J, string) {
	j := &J{P:p}
	VondrakZenith(j)
	VondrakTemperature(j)
	RandVelocity(j)
	ButlerRandDirection(j)
	VondrakPositionJump(j)
	VondrakFlightTime(j)

	msg := CheckLost(j)
	if msg == "" {
		msg = CaptureVondrak(j)
	}

	return j, msg
}

func NewJump(funcs ...JumpMethod) JumpFunc {
	return func(p *P) (*J, string) {
		j := &J{P:p}
		
		for _,f := range funcs {
			msg := f(j)
			if msg != "" {
				return j, msg
			}
		}

		return j,CheckLost(j)
	}
}

func NewJumpSimple(funcs ...interface{}) JumpFunc {
	return NewJump(liftAll(funcs...)...)
}

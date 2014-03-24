package bouncing

import (
	"errors"
	"fmt"
	"flag"
	"strings"
)

var customJumpMethods = map[string]JumpMethod {
	"TemperatureButler": ButlerTemperature, "TemperatureVondrak": VondrakTemperature,
	"RandDirectionMine": RandDirection, "RandDirectionButler": ButlerRandDirection, "RandDirectionVondrak": ButlerRandDirection,
	"PositionJumpButler": ButlerPositionJump, "PositionJumpVondrak": VondrakPositionJump,
	"FlightTimeMine": FlightTime, "FlightTimeButler": VondrakFlightTime, "FlightTimeVondrak": VondrakFlightTime,
	"IsCaptureButler": IsCaptureButler, "IsCaptureVondrak": IsCaptureVondrak,
}

type jumpFlag struct {
	value string
	jump JumpFunc
}

func (f *jumpFlag) String() string {
	return f.value
}

func (f *jumpFlag) Set(value string) error {
	switch value {
	case "Jump":
		f.jump = Jump
	case "Butler":
		f.jump = ButlerJump
	case "Vondrak":
		f.jump = VondrakJump
	default:
		vs := strings.Split(value, ",")
		fn := make([]JumpMethod, len(vs))

		for i, s := range strings.Split(value, ",") {
			f, ok := customJumpMethods[s]
			if !ok {
				return errors.New("invalid JumpMethod flag " + value)
			}
			fn[i] = f
		}

		f.jump = NewJump(fn...)
	}
	return nil
}

var j jumpFlag
var randPositionFuncFlag string
var typ ParticleType

func init() {
	j = jumpFlag{jump:Jump, value:"Jump"}
	flag.Var(&j, "JumpFunc", "jump function to use. One of Butler, Vondrak, Jump, or a commma-separated list of JumpMethods")

	flag.Var(&typ,"ParticleType", "Water, Hydrogen")

	flag.StringVar(&randPositionFuncFlag, "RandInitialPosition", "Butler", "Butler, Vondrak")

	flag.Float64Var(&Tau,"Tau", 6.7e4, "Photodestruction timescale in seconds")
}

func ParseFlags() (JumpFunc, ParticleGenerator) {
	flag.Parse()

	var randPosition func(*P)
	switch randPositionFuncFlag {
	case "Butler":
		randPosition = RandInitialPositionButler
	case "Vondrak":
		randPosition = RandInitialPositionVondrak
	}

	newParticle := func() *P {
		p := &P{Type:typ}
		randPosition(p)
		return p
	}

	PrintFlags()

	return j.jump, newParticle
}

func PrintFlags() {
	flag.VisitAll( func(f *flag.Flag) {
		fmt.Printf("--%v=%v\n",f.Name,f.Value)
	} )
}

// Implement flag.Value interfaces
func (p ParticleType) String() string {
	switch p {
	case Water:
		return "Water"
	case Hydrogen:
		return "Hydrogen"
	}
	return fmt.Sprint(p)
}

func (p *ParticleType) Set(value string) error {
	switch value {
	case "Water":
		*p = Water
	case "Hydrogen":
		*p = Hydrogen
	default:
		return errors.New("invalid ParticleType flag " + value)
	}
	return nil
}

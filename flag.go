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
	jump JumpFunc
	value string
}

func (f *jumpFlag) String() string {
	return f.value
}

func (f *jumpFlag) Set(value string) error {
	switch value {
	case "Jump":
		f.jump = Jump
		f.value = "Jump=ButlerTemperature,RandVelocity,RandDirection,ButlerPositionJump,FlightTime"
	case "Butler":
		f.jump = ButlerJump
	case "Vondrak":
		f.jump = VondrakJump
	case "JumpWithVondrak":
		f.jump = JumpWithVondrak
		f.value = "JumpWithVondrak=VondrakZenith;VondrakTemperature;RandVelocity;RandDirection;ButlerPositionJump;FlightTime;IsLost;IsCaptureVondrak"
	default:
		vs := strings.Split(value, ",")
		fn := make([]JumpMethod, len(vs))

		for i, s := range vs {
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

var particleMutators = map[string]ParticleMutator {
	"RandInitialPositionButler": RandInitialPositionButler,
	"RandInitialPositionVondrak": RandInitialPositionVondrak,
}

type particleFlag struct {
	gen ParticleGenerator
	value string
}

func (f *particleFlag) String() string {
	return f.value
}

func (f *particleFlag) Set(value string) error {
	f.value = value

	switch value {
	case "Butler":
		f.gen = ParticleGeneratorButler
	case "Vondrak":
		f.gen = ParticleGeneratorVondrak
	default:
		vs := strings.Split(value, ",")
		fn := make([]ParticleMutator, len(vs))

		var typ ParticleType
		switch vs[0] {
		case "Water":
			typ = Water
		case "Hydrogen":
			typ = Hydrogen
		}

		vs = vs[1:]

		for i, s := range vs {
			f, ok := particleMutators[s]
			if !ok {
				return errors.New("invalid ParticleMutator " + s)
			}
			fn[i] = f
		}

		f.gen = NewParticleGenerator(typ, fn...)
	}

	return nil
}

var j jumpFlag
var p particleFlag

func init() {
	j = jumpFlag{jump:Jump, value:"Jump"}
	flag.Var(&j, "JumpFunc", "jump function to use. One of Butler, Vondrak, Jump, or a commma-separated list of JumpMethods")

	p = particleFlag{gen:ParticleGeneratorButler, value:"Butler"}
	flag.Var(&p, "ParticleGenerator", "particle generator to use. One of Butler, Vondrak, or a commma-separated list of ParticleType and ParticleMutators")

	flag.Float64Var(&Tau,"Tau", 6.7e4, "Photodestruction timescale in seconds")
}

func PrintFlags() {
	flag.VisitAll( func(f *flag.Flag) {
		fmt.Printf("# --%v=%v\n",f.Name,f.Value)
	} )
}

func ParseFlags() (JumpFunc, ParticleGenerator) {
	flag.Parse()

	PrintFlags()

	return j.jump, p.gen
}

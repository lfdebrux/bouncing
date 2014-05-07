package bouncing

import (
	"errors"
	"fmt"
	"flag"
	"strings"
)

var customJumpMethods = map[string]JumpMethod {
	"TemperatureButler": lift(ButlerTemperature), "TemperatureVondrak": VondrakTemperature,
	"RandVelocity": RandVelocity,
	"RandDirection": lift(RandDirection), "RandDirectionButler": lift(ButlerRandDirection), "RandDirectionVondrak": lift(ButlerRandDirection),
	"PositionJumpButler": lift(ButlerPositionJump), "PositionJumpVondrak": lift(VondrakPositionJump),
	"FlightTime": FlightTime, "FlightTimeButler": VondrakFlightTime, "FlightTimeVondrak": VondrakFlightTime,
	"CheckLost": CheckLost,
	"CaptureButler": CaptureButler, "CaptureVondrak": CaptureVondrak,
}

type jumpFlag struct {
	jump *JumpFunc
	value string
}

func (f *jumpFlag) String() string {
	return f.value
}

func (f *jumpFlag) Set(value string) error {
	switch value {
	case "Jump":
		*f.jump = Jump
		f.value = "Jump=ButlerTemperature,RandVelocity,RandDirection,ButlerPositionJump,FlightTime"
	case "Butler":
		*f.jump = ButlerJump
	case "Vondrak":
		*f.jump = VondrakJump
	case "JumpWithVondrak":
		*f.jump = JumpWithVondrak
		f.value = "JumpWithVondrak=VondrakZenith;VondrakTemperature;RandVelocity;RandDirection;ButlerPositionJump;FlightTime;CheckLost;CaptureVondrak"
	default:
		vs := strings.Split(value, ",")
		fn := make([]JumpMethod, len(vs))

		for i, s := range vs {
			f, ok := customJumpMethods[s]
			if !ok {
				return errors.New("invalid JumpMethod " + s)
			}
			fn[i] = f
		}

		*f.jump = NewJump(fn...)
	}
	return nil
}

var particleMutators = map[string]ParticleMutator {
	"RandInitialPositionButler": RandInitialPositionButler,
	"RandInitialPositionVondrak": RandInitialPositionVondrak,
}

type particleFlag struct {
	gen *ParticleGenerator
	value string
}

func (f *particleFlag) String() string {
	return f.value
}

func (f *particleFlag) Set(value string) error {
	f.value = value

	switch value {
	case "Butler":
		*f.gen = ParticleGeneratorButler
	case "Vondrak":
		*f.gen = ParticleGeneratorVondrak
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

		*f.gen = NewParticleGenerator(typ, fn...)
	}

	return nil
}

func JumpFlagVar(p *JumpFunc, name string, value string, usage string) {
	f := jumpFlag{jump:p, value:value}
	flag.Var(&f, name, usage)
}

func JumpFlag() *JumpFunc {
	var p JumpFunc = Jump
	JumpFlagVar(&p, "JumpFunc", "Jump", "jump function to use. One of Butler, Vondrak, Jump, or a commma-separated list of JumpMethods")
	return &p
}

func ParticleGeneratorFlagVar(p *ParticleGenerator, name string, value string, usage string) {
	f := particleFlag{gen:p, value:value}
	flag.Var(&f, name, usage)
}

func ParticleGeneratorFlag() *ParticleGenerator {
	var p ParticleGenerator = ParticleGeneratorButler
	ParticleGeneratorFlagVar(&p, "ParticleGenerator", "Butler", "particle generator to use. One of Butler, Vondrak, or a commma-separated list of ParticleType and ParticleMutators")
	return &p
}
// func init() {
// 	j = jumpFlag{jump:Jump, value:"Jump"}
// 	flag.Var(&j, "JumpFunc", "jump function to use. One of Butler, Vondrak, Jump, or a commma-separated list of JumpMethods")

// 	p = particleFlag{gen:ParticleGeneratorButler, value:"Butler"}
// 	flag.Var(&p, "ParticleGenerator", "particle generator to use. One of Butler, Vondrak, or a commma-separated list of ParticleType and ParticleMutators")

// 	flag.Float64Var(&Tau,"Tau", 6.7e4, "Photodestruction timescale in seconds")
// }

func PrintFlag(f *flag.Flag) {
	fmt.Printf("# -%v=%v\n",f.Name,f.Value)
}

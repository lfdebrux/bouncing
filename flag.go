package bouncing

import (
	"errors"
	"flag"
	"fmt"
)

var customJumpFlags flag.FlagSet

var customJumpMethodFlags = map[string]map[string]JumpMethod {
	"Temperature": {"Butler": ButlerTemperature, "Vondrak": VondrakTemperature},
	"RandDirection": {"Mine": RandDirection, "Butler": ButlerRandDirection, "Vondrak": ButlerRandDirection},
	"PositionJump": {"Butler": ButlerPositionJump, "Vondrak": VondrakPositionJump},
	"FlightTime": {"Mine": FlightTime, "Butler": VondrakFlightTime, "Vondrak": VondrakFlightTime},
	"IsCapture": {"Butler": IsCaptureButler, "Vondrak": IsCaptureVondrak},
}

var jumpFuncFlag string
var randPositionFuncFlag string
var particletype ParticleType

func init() {
	flag.StringVar(&jumpFuncFlag, "JumpFunc", "Jump", "Butler, Vondrak, Jump, Custom")

	flag.Var(&particletype,"ParticleType", "Water, Hydrogen")

	flag.StringVar(&randPositionFuncFlag, "RandInitialPosition", "Butler", "Butler, Vondrak")

	flag.Float64Var(&Tau,"Tau", 6.7e4, "Photodestruction timescale in seconds")

	customJumpFlags := flag.NewFlagSet("Custom JumpFunc", flag.ExitOnError)

	customJumpFlags.String("Temperature", "Butler", "Butler, Vondrak")
	customJumpFlags.String("RandDirection", "Mine", "Mine, Butler")
	customJumpFlags.String("PositionJump", "Butler", "Butler, Vondrak")
	customJumpFlags.String("FlightTime", "Mine", "Mine, Butler, Vondrak")
	customJumpFlags.String("IsCapture", "Butler", "Butler, Vondrak")
}

func ParseFlags() (JumpFunc, func() *P) {
	flag.Parse()

	var jump JumpFunc
	switch jumpFuncFlag {
	case "Custom":
		jump = parseCustomJumpFlags(flag.Args())
	case "Jump":
		jump = Jump
	case "Butler":
		jump = ButlerJump
	case "Vondrak":
		jump = VondrakJump
	}

	var randPosition func(*P)
	switch randPositionFuncFlag {
	case "Butler":
		randPosition = RandInitialPositionButler
	case "Vondrak":
		randPosition = RandInitialPositionVondrak
	}

	newParticle := func() *P {
		p := &P{Type:particletype}
		randPosition(p)
		return p
	}

	PrintFlags()

	return jump, newParticle
}

func PrintFlags() {
	printFlags := func(f *flag.Flag) {
		fmt.Printf("--%v=%v\n",f.Name,f.Value)
	}

	flag.VisitAll(printFlags)

	if customJumpFlags.Parsed() {
		customJumpFlags.VisitAll(printFlags)
	}
}

func customJumpMethodFlagsGet(s flag.FlagSet, name string) JumpMethod {
	f := s.Lookup(name)
	fn, ok := customJumpMethodFlags[f.Name][f.Value.String()]
	if !ok {
		panic("invalid JumpMethod flag")
	}
	return fn
}

func parseCustomJumpFlags(args []string) JumpFunc {
	customJumpFlags.Parse(args)

	tm := customJumpMethodFlagsGet(customJumpFlags, "Temperature")
	rd := customJumpMethodFlagsGet(customJumpFlags, "RandDirection")
	pj := customJumpMethodFlagsGet(customJumpFlags, "PositionJump")
	ft := customJumpMethodFlagsGet(customJumpFlags, "FlightTime")
	cp := customJumpMethodFlagsGet(customJumpFlags, "IsCapture")

	return NewJump(tm,RandVelocity,rd,pj,ft,cp)
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
		return errors.New("invalid ParticleType flag value")
	}
	return nil
}

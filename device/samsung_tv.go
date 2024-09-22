package device

import (
	"fmt"
)

type Tv struct {
	OnOff
	CommandRunner
	id  string
	hub *HarmonyHub
}

func (tv *Tv) RunCommand(cmd string) (any, error) {
	switch cmd {
	case "on":
		return nil, tv.On()
	case "off":
		return nil, tv.Off()
	default:
		return nil, fmt.Errorf("unknown command %v for Tv", cmd)
	}
}

func (tv *Tv) On() (err error) {
	return tv.hub.SendCommand("PowerOn", tv.id)
}
func (tv *Tv) Off() (err error) {
	return tv.hub.SendCommand("PowerOff", tv.id)
}

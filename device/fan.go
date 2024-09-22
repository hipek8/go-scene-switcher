package device

import "fmt"

type Fan struct {
	CommandRunner
	id  string
	hub *HarmonyHub
}

func (fan *Fan) RunCommand(cmd string) (any, error) {
	switch cmd {
	case "toggle", "on", "off":
		return nil, fan.Toggle()
	case "oscillate", "osc", "turn", "move":
		return nil, fan.ToggleOscillation()
	case "up", "+":
		return nil, fan.SpeedUp()
	case "down", "-":
		return nil, fan.SpeedDown()
	default:
		return nil, fmt.Errorf("unknown command %v for Fan", cmd)
	}
}
func (fan *Fan) Toggle() error {
	return fan.hub.SendCommand("PowerToggle", fan.id)
}

func (fan *Fan) SpeedUp() error {
	return fan.hub.SendCommand("+ Speed", fan.id)
}

func (fan *Fan) SpeedDown() error {
	return fan.hub.SendCommand("- Speed", fan.id)
}

func (fan *Fan) ToggleOscillation() error {
	return fan.hub.SendCommand("Oscillation", fan.id)
}

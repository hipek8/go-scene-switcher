package device

import "time"

type SwitchOn interface {
	On() (err error)
}
type SwitchOff interface {
	Off() (err error)
}
type OnOff interface {
	SwitchOn
	SwitchOff
}

type DeviceStatus int

const (
	OFF DeviceStatus = iota
	ON
	STANDBY
	UNKNOWN
)

type TV struct {
	isOn DeviceStatus
}

func (tv *TV) On() {
	time.Sleep(time.Duration(2 * time.Second))
	tv.isOn = ON
}

func (tv *TV) Off() {
	tv.isOn = OFF
}

func (tv *TV) Status() (status DeviceStatus) {
	return tv.isOn
}

func NewTV() *TV {
	return &TV{isOn: UNKNOWN}
}

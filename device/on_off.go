package device

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

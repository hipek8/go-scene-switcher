package device_test

import (
	"testing"

	"my/scene-switcher/device"
)

func TestTVOn(t *testing.T) {
	tv := device.NewTV()
	tv.On()
	if tv.Status() != device.ON {
		t.Fail()
	}
}

func TestTVOff(t *testing.T) {
	tv := device.NewTV()
	tv.Off()
	if tv.Status() != device.OFF {
		t.Fail()
	}
}

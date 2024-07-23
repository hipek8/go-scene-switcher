package device_test

import (
	"testing"

	"my/scene-switcher/device"
)

func TestWsUrl(t *testing.T) {
	hub := device.NewHarmonyHub("1.2.3.4", 4321, "abc")
	if hub.WsUrl() != "ws://1.2.3.4:4321/?domain=svcs.myharmony.com&hubId=abc" {
		t.Fail()
	}
}

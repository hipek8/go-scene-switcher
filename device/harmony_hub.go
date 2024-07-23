package device

import (
	"encoding/json"
	"fmt"
	"time"
)

// CONFIG
const HOST = "192.168.0.17"
const PORT = 8088
const HUB_ID = "16366653"
const TV_ID = "69850160"

var i int = 1

type tv struct {
	OnOff
	id  string
	hub *HarmonyHub
}

func (tv *tv) On() (err error) {
	return tv.hub.SendCommand("PowerOn", tv.id)
}
func (tv *tv) Off() (err error) {
	return tv.hub.SendCommand("PowerOff", tv.id)
}

type HarmonyHub struct {
	host   string
	port   int
	hub_id string
}

func (hub *HarmonyHub) WsUrl() string {
	return fmt.Sprintf(
		"ws://%v:%v/?domain=svcs.myharmony.com&hubId=%v",
		hub.host,
		hub.port,
		hub.hub_id)
}

func (hub *HarmonyHub) GetTv() OnOff {
	return &tv{id: TV_ID, hub: hub}
}

func (hub *HarmonyHub) SendCommand(cmd string, dev_id string) (err error) {
	// TODO: ws connect

	action, err := json.Marshal(
		map[string]interface{}{
			"command":  cmd,
			"type":     "IRCommand",
			"deviceId": dev_id,
		},
	)
	if err != nil {
		return err
	}
	i++
	dumped, err := json.Marshal(map[string]interface{}{
		"hubId":   hub.hub_id,
		"timeout": 10,
		"Hbus": map[string]interface{}{
			"command": cmd,
			"id":      i,
			"params": map[string]interface{}{
				"status":    "press",
				"verb":      "render",
				"timestamp": time.Now().Unix(),
				"action":    string(action),
			},
		},
	})
	if err != nil {
		return
	}
	fmt.Println(string(dumped))
	return
}
func (hub *HarmonyHub) StartActivity(activity_id string) {
	// TODO: ws connect
}

func MyHarmonyHub() *HarmonyHub {
	return &HarmonyHub{
		host:   HOST,
		port:   PORT,
		hub_id: HUB_ID,
	}
}

func NewHarmonyHub(host string, port int, hub_id string) *HarmonyHub {
	return &HarmonyHub{
		host:   host,
		port:   port,
		hub_id: hub_id,
	}
}

func ExampleHarmonyHubUsage() {
	hub := MyHarmonyHub()
	tv := hub.GetTv()
	tv.On()
}

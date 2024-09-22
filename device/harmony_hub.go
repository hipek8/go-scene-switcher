package device

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

// CONFIG
const HOST = "192.168.0.17"
const PORT = 8088
const HUB_ID = "16366653"
const TV_ID = "69850160"
const FAN_ID = "70675878"

var i int = 1

type HarmonyHub struct {
	CommandRunner
	host   string
	port   int
	hub_id string
}

func (hub *HarmonyHub) RunCommand(cmd string) (any, error) {
	switch cmd {
	default:
		return nil, fmt.Errorf("unknown command %v for HarmonyHub", cmd)
	}
}

func (hub *HarmonyHub) WsUrl() string {
	return fmt.Sprintf(
		"ws://%v:%v/?domain=svcs.myharmony.com&hubId=%v",
		hub.host,
		hub.port,
		hub.hub_id)
}

func (hub *HarmonyHub) GetTv() *Tv {
	return &Tv{id: TV_ID, hub: hub}
}

func (hub *HarmonyHub) GetFan() *Fan {
	return &Fan{id: FAN_ID, hub: hub}
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
		"hbus": map[string]interface{}{
			"cmd": "vnd.logitech.harmony/vnd.logitech.harmony.engine?holdAction",
			"id":  i,
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
	err = hub.sendWs(hub.WsUrl(), dumped)
	return err
}

func (hub *HarmonyHub) sendWs(url string, data []byte) error {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}
	return nil
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

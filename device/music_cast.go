package device

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MusicCast struct {
	OnOff
	CommandRunner
	host string
	zone string
}

func (mc *MusicCast) RunCommand(cmd string) (any, error) {
	switch cmd {
	case "on":
		return nil, mc.On()
	case "off":
		return nil, mc.Off()
	case "pc":
		return nil, mc.SetInput("pc")
	case "net_radio":
		return nil, mc.SetInput("net_radio")
	case "rns":
		return nil, mc.RunPreset("1")
	case "r357":
		return nil, mc.RunPreset("2")
	case "vol/low": // rns
		return nil, mc.SetVolume("60")
	case "vol/mid": // r357
		return nil, mc.SetVolume("90")
	case "vol/high": // for PC
		return nil, mc.SetVolume("120")
	case "vol/xxl": // for Mixer/AudioInterface
		return nil, mc.SetVolume("136")
	default:
		return nil, fmt.Errorf("unknown command %v for MusicCast", cmd)
	}
}

var STATUS_MAP map[string]DeviceStatus = map[string]DeviceStatus{
	"on":      ON,
	"standby": STANDBY,
	"off":     OFF,
}

func (status *DeviceStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	st, ok := STATUS_MAP[s]
	if !ok {
		return fmt.Errorf("unknown status")
	}
	*status = st
	return nil
}

type MusicCastStatus struct {
	Volume    int          `json:"volume"`
	Power     DeviceStatus `json:"power"`
	Mute      bool         `json:"mute"`
	InputText string       `json:"input_text"`
}

func (mc *MusicCast) MainUrl(endpoint string) string {
	return fmt.Sprintf(
		"http://%v/YamahaExtendedControl/v1/%v",
		mc.host,
		endpoint)
}

func (mc *MusicCast) ZoneUrl(endpoint string) string {
	return mc.MainUrl(fmt.Sprintf("%v/%v", mc.zone, endpoint))
}

func (mc *MusicCast) NetUsbUrl(endpoint string) string {
	return mc.MainUrl(fmt.Sprintf("netusb/%v", endpoint))
}

func (mc *MusicCast) MakeGet(url string) ([]byte, error) {
	fmt.Println("GET: " + url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad response: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (mc *MusicCast) GetStatus() (*MusicCastStatus, error) {
	response, err := mc.MakeGet(mc.ZoneUrl("getStatus"))
	if err != nil {
		return nil, err
	}
	var status MusicCastStatus
	err = json.Unmarshal(response, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (mc *MusicCast) On() error {
	_, err := mc.MakeGet(mc.ZoneUrl("setPower?power=on"))
	return err
}

func (mc *MusicCast) Off() error {
	_, err := mc.MakeGet(mc.ZoneUrl("setPower?power=standby"))
	return err
}

func (mc *MusicCast) Mute() error {
	_, err := mc.MakeGet(mc.ZoneUrl("setMute?enable=true"))
	return err
}

func (mc *MusicCast) Unmute() error {
	_, err := mc.MakeGet(mc.ZoneUrl("setMute?enable=false"))
	return err
}

func (mc *MusicCast) SetInput(input string) error {
	_, err := mc.MakeGet(mc.ZoneUrl("setInput?input=" + input))
	return err
}

func (mc *MusicCast) SetVolume(volume string) error {
	_, err := mc.MakeGet(mc.ZoneUrl("setVolume?volume=" + volume))
	return err
}

func (mc *MusicCast) SetSoundProgram(program string) error {
	_, err := mc.MakeGet(mc.ZoneUrl("setSoundProgram?program=" + program))
	return err
}

func (mc *MusicCast) RunPreset(presetNumber string) error {
	// notice presentNumber is 1-indexed, i.e. the first element from netusb/getPresetInfo is 1
	_, err := mc.MakeGet(mc.NetUsbUrl(fmt.Sprintf("recallPreset?zone=%v&num=%v", mc.zone, presetNumber)))
	return err
}

func (mc *MusicCast) SetScene(
	input string,
	volume string,
	preset string,
) error {
	defaultProgram := "straight"
	if err := mc.On(); err != nil {
		return err
	}
	if err := mc.Mute(); err != nil {
		return err
	}
	defer mc.Unmute()
	if err := mc.SetVolume(volume); err != nil {
		return err
	}
	if err := mc.SetInput(input); err != nil {
		return err
	}
	if preset != "" {
		err := mc.RunPreset(preset)
		if err != nil {
			return err
		}
	}
	if err := mc.SetSoundProgram(defaultProgram); err != nil {
		return err
	}
	return nil
}

func MyMusicCast() *MusicCast {
	return &MusicCast{host: "192.168.1.104", zone: "main"}
}

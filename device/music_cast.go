package device

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MusicCast struct {
	OnOff
	host string
	zone string
}

func MyMusicCast() MusicCast {
	return MusicCast{host: "1.2.3.4", zone: "other"}
}

var STATUS_MAP map[string]DeviceStatus = map[string]DeviceStatus{
	"on":      ON,
	"standby": STANDBY,
	"off":     OFF}

type MusicCastStatus struct {
	volume int
	power  string
}

func (mc *MusicCast) MainUrl(endpoint string) string {
	return fmt.Sprintf(
		"http://%v/YamahaExtendedControl/v1/%v/%v",
		mc.host,
		mc.zone,
		endpoint)
}

func (mc *MusicCast) ZoneRequestGet(endpoint string) (interface{}, error) {
	url := mc.MainUrl(endpoint)
	fmt.Println("GET: %v", url)
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
	var jsonResponse map[string]interface{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}

func (mc *MusicCast) GetStatus() (DeviceStatus, error) {
	return STATUS_MAP["standby"], nil
}

func (mc *MusicCast) On() error {
	_, err := mc.ZoneRequestGet("setPower?power=on")
	return err
}

func (mc *MusicCast) Off() error {
	_, err := mc.ZoneRequestGet("setPower?power=standby")
	return err
}

func (mc *MusicCast) Mute() error {
	_, err := mc.ZoneRequestGet("setMute?enable=true")
	return err
}

func (mc *MusicCast) Unmute() error {
	_, err := mc.ZoneRequestGet("setMute?enable=false")
	return err
}

func (mc *MusicCast) SetInput(input string) error {
	_, err := mc.ZoneRequestGet("setInput?input=" + input)
	return err
}

func (mc *MusicCast) SetVolume(volume string) error {
	_, err := mc.ZoneRequestGet("setVolume?volume=" + volume)
	return err
}

func (mc *MusicCast) SetSoundProgram(program string) error {
	_, err := mc.ZoneRequestGet("setSoundProgram?program=" + program)
	return err
}
func (mc *MusicCast) SetScene(
	input string,
	soundProgram string,
	volume int) error {
	if err := mc.On(); err != nil {
		return err
	}
	if err := mc.Mute(); err != nil {
		return err
	}
	defer mc.Unmute()
	if err := mc.SetVolume(string(volume)); err != nil {
		return err
	}
	if err := mc.SetInput(input); err != nil {
		return err
	}
	if err := mc.SetSoundProgram(soundProgram); err != nil {
		return err
	}
	return nil
}

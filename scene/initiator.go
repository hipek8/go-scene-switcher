package scene

import (
	"fmt"
	"my/scene-switcher/device"
	"strings"
	"time"
)

type SceneScheduler interface {
	SetScene(scene string) error
}

type BaseSceneScheduler struct {
	Set chan<- string
}

func (setter *BaseSceneScheduler) SetScene(scene string) error {
	setter.Set <- scene
	return nil
}

// Sets scene when requested via API
type ApiSceneScheduler struct {
	BaseSceneScheduler
}

// Sets scene when observing change in MusicCast scene
type MusicCastSceneScheduler struct {
	BaseSceneScheduler
}

func (mci *MusicCastSceneScheduler) Run() error {
	mc := device.MyMusicCast()
	tv := device.MyHarmonyHub().GetTv()
	status, err := mc.GetStatus()
	if err != nil {
		fmt.Printf("Error getting MusicCast status %v", err)
	}
	var lastStatus *device.MusicCastStatus = status
	for {
		time.Sleep(time.Duration(time.Second * 5))
		status, err := mc.GetStatus()
		fmt.Printf("Power: %v, Input: %v\n", status.Power, status.InputText)
		if err != nil {
			fmt.Printf("Error getting MusicCast status %v\n", err)
		}
		if status.Power != lastStatus.Power || status.InputText != lastStatus.InputText {
			lastStatus = status
			if status.Power == device.STANDBY {
				fmt.Printf("Turning TV Off\n")
				tv.Off()
			} else {
				if strings.ToLower(status.InputText) == "pc" && status.Power == device.ON {
					fmt.Printf("Turning TV On\n")
					tv.On()
				}
			}
		}
	}
}

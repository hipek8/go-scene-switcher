package scene

import (
	"fmt"
	"my/scene-switcher/device"
)

type Synchronizer interface {
	Run()
	Stop()
}

type BaseSynchronizer struct {
	Sync         chan string
	currentScene string
	isRunning    bool
}

type DummySynchronizer struct {
	BaseSynchronizer
}

func (sync *DummySynchronizer) Stop() {
	sync.Sync <- ""
}

func (sync *DummySynchronizer) CurrentScene() string {
	return sync.currentScene
}

func (sync *DummySynchronizer) SetScene(scene string, mc *device.MusicCast, tv *device.Tv) {
	if mc == nil {
		mc = device.MyMusicCast()
	}
	if tv == nil {
		tv = device.MyHarmonyHub().GetTv()
	}
	switch scene {
	case "off":
		{
			go tv.Off()
			go mc.Off()
		}
	case "pc":
		{
			go mc.SetScene("hdmi1", "120", "")
			go tv.On()
		}
	case "rpi", "kodi":
		{
			go tv.On()
			go mc.SetScene("rpi", "120", "")
		}
	case "rns":
		{
			go mc.SetScene("netusb", "80", "1")
			go tv.Off()
		}
	case "r357":
		{
			go mc.SetScene("netusb", "95", "2")
			go tv.Off()
		}
	case "tv":
		{
			go tv.On()
			go mc.SetScene("hdmi1", "100", "")
		}
	case "ai", "interface", "int":
		{
			go tv.Off()
			go mc.SetScene("av2", "136", "")
		}
	case "mixer":
		{
			go tv.Off()
			go mc.SetScene("av3", "136", "")
		}
	}
}

func (sync *DummySynchronizer) Run() {
	mc := device.MyMusicCast()
	tv := device.MyHarmonyHub().GetTv()
	if sync.isRunning {
		return
	}
	sync.isRunning = true
	defer func() { sync.isRunning = false }()
	for scene := range sync.Sync {
		if scene == "" {
			break
		}
		sync.currentScene = scene
		fmt.Println("Synchronizing scene: " + scene)
		sync.SetScene(scene, mc, tv)
	}
}

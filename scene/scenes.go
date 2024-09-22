package scene

import (
	"fmt"
	"my/scene-switcher/device"
)

func SetScene(name string, tv *device.Tv, mc *device.MusicCast) error {
	if name == "on" {
		err := tv.On()
		fmt.Println("MC On")
		// err := mc.On()
		if err != nil {
			fmt.Println(err.Error())
		}
	} else if name == "off" {
		fmt.Println("MC Off")
		err := tv.Off()
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Unknown scene")
	}
	return nil
}

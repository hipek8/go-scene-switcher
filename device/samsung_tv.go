package device

import (
	"fmt"
	"time"
)

type SamsungTv struct {
	OnOff
	port     int
	mac      string
	hostname string
}

func (tv *SamsungTv) On() error {
	fmt.Println("On: " + string(tv.port) + tv.mac + tv.hostname)
	harmony_tv := MyHarmonyHub().GetTv()
	for retry := 1; retry <= 3; retry++ {
		err := harmony_tv.On()
		if err != nil {
			// TODO websocket on
			_ = err
		}
		time.Sleep(time.Second * time.Duration(retry))
		// TODO check status and break

	}
	return nil
}

func (tv *SamsungTv) Off() error {
	fmt.Println("Off: " + string(tv.port) + tv.mac + tv.hostname)
	harmony_tv := MyHarmonyHub().GetTv()
	for retry := 0; retry < 3; retry++ {
		err := harmony_tv.Off()
		if err != nil {
			// TODO websocket on
			_ = err
		}
		time.Sleep(time.Second * 1)
		// TODO check status and break

	}
	return nil
}

package device

import (
	"fmt"
)

type Pool struct {
	deviceMap map[string]CommandRunner
}

func (pool *Pool) Get(deviceName string) (CommandRunner, error) {
	dev, ok := pool.deviceMap[deviceName]
	if ok {
		return dev, nil
	} else {
		return nil, fmt.Errorf("couldn't find device %v", deviceName)
	}
}

func MyPool() *Pool {
	hh := MyHarmonyHub()
	return &Pool{
		deviceMap: map[string]CommandRunner{
			"mc":  MyMusicCast(),
			"hh":  hh,
			"tv":  hh.GetTv(),
			"fan": hh.GetFan(),
		},
	}
}

package scene

import "fmt"

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
func (sync *DummySynchronizer) Run() {
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
	}
}

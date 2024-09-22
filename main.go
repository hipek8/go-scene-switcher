package main

import (
	"my/scene-switcher/api"
	"my/scene-switcher/scene"
)

func main() {
	schedulerToSynchronizerChannel := make(chan string)
	syncer := scene.DummySynchronizer{
		BaseSynchronizer: scene.BaseSynchronizer{
			Sync: schedulerToSynchronizerChannel}}
	go syncer.Run()

	r := SetupRouter()
	r = api.SetupSceneEndpoint(r, schedulerToSynchronizerChannel)
	api.SetupDeviceEndpoint(r)
	mcScheduler := scene.MusicCastSceneScheduler{}
	mcScheduler.Run()
	r.Run(":8080")

}

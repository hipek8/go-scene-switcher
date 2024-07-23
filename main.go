package main

import (
	"my/scene-switcher/api"
	"my/scene-switcher/scene"
)

func main() {
	messageChannel := make(chan string)
	syncer := scene.DummySynchronizer{
		BaseSynchronizer: scene.BaseSynchronizer{
			Sync: messageChannel}}
	go syncer.Run()

	r := SetupRouter()
	r = api.SetupSceneEndpoint(r, messageChannel)
	r.Run(":8080")
}

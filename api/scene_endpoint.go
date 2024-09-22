package api

import (
	"my/scene-switcher/scene"

	"github.com/gin-gonic/gin"
)

func SetupSceneEndpoint(r *gin.Engine, messageChannel chan string) *gin.Engine {
	setter1 := scene.ApiSceneScheduler{
		BaseSceneScheduler: scene.BaseSceneScheduler{
			Set: messageChannel}}
	r.POST("/scene/:scene", func(c *gin.Context) {
		sceneName := c.Param("scene")
		setter1.SetScene(sceneName)
	})
	return r
}

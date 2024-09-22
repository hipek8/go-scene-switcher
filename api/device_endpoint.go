package api

import (
	"fmt"
	"my/scene-switcher/device"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupDeviceEndpoint(r *gin.Engine) *gin.Engine {
	pool := device.MyPool()

	r.POST("/device/:device/:command", func(c *gin.Context) {
		deviceName := c.Param("device")
		command := c.Param("command")
		device_, err := pool.Get(deviceName)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		response, err := device_.RunCommand(command)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		} else {
			c.String(http.StatusOK, fmt.Sprintf("ok: %v", response))
		}
	})
	return r
}

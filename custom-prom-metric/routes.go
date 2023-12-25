package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRoutes() {
	client := ConnectToRedis()
	r := gin.Default()
	// GET /devices
	r.GET("/devices", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetDevices(client, c.Writer, c.Request))
	})
	// GET /devices/:id
	r.GET("/devices/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetDevice(client, c.Param("id"), c))
	})
	// POST /devices
	r.POST("/devices", func(c *gin.Context) {
		var device Device
		c.BindJSON(&device)
		StoreDevice(client, device, c)
	})
	// DELETE /devices/:id
	r.DELETE("/devices/:id", func(c *gin.Context) {
		DeleteDevice(client, c.Param("id"), c)
	})
	// GET /metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// Run the server
	r.Run(":8080")
}

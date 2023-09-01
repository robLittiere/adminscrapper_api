package controllers

import (
	"adminscrapper/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.RouterGroup) {
	deviceGroup := r.Group("/devices")
	{
		deviceGroup.GET("/", GetDevices)

	}
}

func GetDevices(c *gin.Context) {
	// Here parse data from file
	data, err := services.ParseJsonFile("./data/network.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

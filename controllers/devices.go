package controllers

import (
	"adminscrapper/api/services"
	"net/http"
	"os"

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
	datapath := os.Getenv("DATA_FILE")
	data, err := services.ParseJsonFile(datapath)
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

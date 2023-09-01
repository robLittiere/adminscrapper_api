package main

import (
	"adminscrapper/api/controllers"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Init config
	// env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	r := InitRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Enable CORS
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Create "api" group so that we can use "api" as prefix
	// exemple : http://localhost:8080/api/...
	api := r.Group("/api")

	// Create "v1" group so that we can use "v1" as prefix
	// exemple : http://localhost:8080/api/v1/...
	v := api.Group("/v1")

	// Add routes or even group routes to this v group
	// exemple : http://localhost:8080/api/v1/pings
	pings := v.Group("/pings")
	{
		pings.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "We are in the ping group BABY",
			})
		})
	}

	// Group routes
	controllers.AdminRoutes(v)

	return r
}

package main

import (
	"DxServer/database"
	"DxServer/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitMongo("mongodb://localhost:27017/", "DxServer")
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	api := r.Group("/api")
	routers.LoginRouter(api)

	err := r.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}

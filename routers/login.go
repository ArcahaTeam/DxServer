package routers

import (
	"DxServer/services"
	"github.com/gin-gonic/gin"
)

func LoginRouter(rg *gin.RouterGroup) {
	rg.POST("/login", func(c *gin.Context) {
		var json struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"success": false, "error": "Invalid request"})
			return
		}

		success, token := services.Login(json.Username, json.Password)
		if success {
			c.JSON(200, gin.H{"success": true, "token": token})
		} else {
			c.JSON(401, gin.H{"success": false})
		}
	})
}

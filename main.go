package main

import (
	"DxServer/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitMongo("mongodb://localhost:27017/", "DxServer")
	r := gin.Default()

	// 定义一个GET请求路由
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// 定义一个POST请求路由
	r.POST("/login", func(c *gin.Context) {
		var json struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// 绑定请求数据
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": "Bad request"})
			return
		}

		// 验证用户名和密码
		if json.Username == "admin" && json.Password == "password" {
			c.JSON(200, gin.H{
				"message": "Login successful!",
			})
		} else {
			c.JSON(401, gin.H{
				"message": "Invalid username or password",
			})
		}
	})

	// 启动服务器，监听在8080端口
	err := r.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mywebapp/handlers"
)

func main() {
	router := gin.Default()

	// 路由配置
	router.GET("/", homeHandler)
	router.GET("/users", handlers.GetUsers)
	router.POST("/users", handlers.CreateUser)

	// 启动服务器
	if err := router.Run(":8080"); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}

func homeHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome"})
}

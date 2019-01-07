package main

import (
	"rtnews/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//register
	r.POST("/register", controllers.RegisterController)
	r.POST("/login", controllers.LoginController)
	r.POST("/cancel", nil)
	r.GET("/news/:newsId", nil)
	r.Run(":8080")
}

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
	r.GET("/news", controllers.NewsListController)
	r.GET("/user", controllers.UserListController)
	r.POST("/newsUpdate", controllers.NewsUpdateController)
	r.GET("/retriveComment", controllers.RetriveCommentController)
	r.POST("/commentUpdate", controllers.CommentUpdateController)
	r.Run(":8080")
}

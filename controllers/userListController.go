package controllers

import (
	"fmt"
	"rtnews/services"

	"github.com/gin-gonic/gin"
)

func UserListController(ctx *gin.Context) {
	usernameVal, _ := ctx.GetQuery("username")
	data, err := services.UserListService(fmt.Sprintf("username='%s'", usernameVal))
	if err != nil {
		ctx.JSON(403, map[string]interface{}{
			"Code":    1001,
			"Message": "failed",
			"Data":    nil,
		})
		return
	}
	ctx.JSON(200, map[string]interface{}{
		"Code":    1000,
		"Message": "success",
		"Data":    data,
	})
	return
}

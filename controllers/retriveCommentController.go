package controllers

import (
	"fmt"
	"log"
	"rtnews/services"

	"github.com/gin-gonic/gin"
)

func RetriveCommentController(ctx *gin.Context) {
	newIdVal, _ := ctx.GetQuery("newsId")
	data, err := services.RetriveCommentService(fmt.Sprintf("newsId= %s", newIdVal))
	if err != nil {
		log.Print("Call RetriveCommentService fatal error ", err)
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

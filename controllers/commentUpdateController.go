package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"rtnews/services"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Content  string
	AuthorId int
	NewsId   int
}

func CommentUpdateController(ctx *gin.Context) {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	defer ctx.Request.Body.Close()
	if err != nil {
		log.Print("read request body fatal error ", err)
		ctx.JSON(401, map[string]interface{}{
			"Code":    2001,
			"Message": "failed",
			"Data":    nil,
		})
		return
	}
	req := &Request{}
	err = json.Unmarshal(data, req)
	if err != nil {
		log.Print("json unmarshal fatal error ", err)
		ctx.JSON(403, map[string]interface{}{
			"Code":    2002,
			"Message": "failed",
			"Data":    nil,
		})
		return
	}
	log.Print("receive data is <'%s','%s','%s'>", req.Content, req.AuthorId, req.NewsId)
	result, err := services.CommentUpdate(req.Content, req.AuthorId, req.NewsId)
	if err != nil {
		log.Print("call commentUpdateService fatal error ", err)
		ctx.JSON(200, map[string]interface{}{
			"Code":    1001,
			"Message": "failed",
			"Data":    nil,
		})
		return
	}
	if result == true {
		ctx.JSON(200, map[string]interface{}{
			"Code":    1000,
			"Message": "success",
			"Data":    true,
		})
		return
	}
	ctx.JSON(200, map[string]interface{}{
		"Code":    1002,
		"Message": "failed",
		"Data":    false,
	})
	return
}

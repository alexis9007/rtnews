package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"rtnews/services"
	"time"

	"github.com/gin-gonic/gin"
)

type NewsUpdate struct {
	Title       string
	Content     string
	PublishedAt time.Time
	Author      string
}

func NewsUpdateController(ctx *gin.Context) {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	defer ctx.Request.Body.Close()
	if err != nil {
		log.Println("ReadAll fatal error", err)
		ctx.JSON(401, map[string]interface{}{
			"Code":    2001,
			"Message": "ReadAll fatal error",
			"data":    nil,
		})
		return
	}
	upd := &NewsUpdate{
		PublishedAt: time.Now(),
	}
	err = json.Unmarshal(data, upd)
	if err != nil {
		log.Println("json Unmarshal fatal error ", err)
		ctx.JSON(403, map[string]interface{}{
			"Code":    2001,
			"Message": "Wrong data format",
			"Data":    nil,
		})
		return
	}
	log.Printf("receive NewsUpdate data is <%s, %s, %s, %s>", upd.Title, upd.Content, upd.PublishedAt, upd.Author)
	result, err := services.NewsUpdateService(upd.Title, upd.Content, upd.PublishedAt, upd.Author)
	if err != nil {
		log.Println("call NewsUpdateService fatal error ", err)
		ctx.JSON(200, map[string]interface{}{
			"Code":    1001,
			"Message": "NewsUpdate failed",
			"Data":    nil,
		})
		return
	}
	ctx.JSON(200, map[string]interface{}{
		"Code":    200,
		"Message": "NewsUpdate success",
		"Data":    result,
	})

}

package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"rtnews/services"
	"time"

	"github.com/gin-gonic/gin"
)

// UserRequest 代表用户请求
type UserRequest struct {
	Username string
	Password string
	Birthday time.Time
}

// RegisterController 注册接口
func RegisterController(ctx *gin.Context) {
	//ctx.Request.Body
	//reader: can read
	//writer: can write
	data, err := ioutil.ReadAll(ctx.Request.Body)
	defer ctx.Request.Body.Close()
	if err != nil {
		log.Println("ReadAll fatal error", err)
		ctx.JSON(401, map[string]interface{}{
			"Code":    2001,
			"Message": "ReadAll fatal error",
			"Data":    nil,
		})
		return
	}
	req := &UserRequest{} //
	err = json.Unmarshal(data, req)
	if err != nil {
		log.Println("json Unmarshall fatal error", err)
		ctx.JSON(403, map[string]interface{}{
			"Code":    2002,
			"Message": "wrong data format",
			"Data":    nil,
		})
		return
	}
	log.Printf("receive user data <%s,%s>", req.Username, req.Password)

	//call registerService ...
	//registerResult:=services.RegisterService(req.Username,req.Password)
	result, err := services.RegisterService(req.Username, req.Password)
	if !result {
		ctx.JSON(200, map[string]interface{}{
			"Code":    1001,
			"Message": "Register failed",
			"Data":    false,
		})
		return
	}
	ctx.JSON(200, map[string]interface{}{
		"Code":    1000,
		"Message": "Success",
		"Data":    true,
	})
	return
}

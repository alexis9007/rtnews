package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"rtnews/services"

	"github.com/gin-gonic/gin"
)

// LoginController 注册接口
func LoginController(ctx *gin.Context) {
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
	req := &UserRequest{}
	err = json.Unmarshal(data, req)
	if err != nil {
		log.Println("json Unmarshal fatal error", err)
		ctx.JSON(403, map[string]interface{}{
			"Code":    2002,
			"Message": "wrong data",
			"Data":    nil,
		})
		return
	}
	log.Printf("receive user data <'%s', '%s'>", req.Username, req.Password)

	//接入service
	result, err := services.LoginService(req.Username, req.Password)
	if err != nil {
		log.Println("Registe failed", err)
		ctx.JSON(200, map[string]interface{}{
			"Code":    1001,
			"Message": "login failed",
			"Data":    false,
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
}

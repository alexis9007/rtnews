package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"rtnews/services"
	"time"

	"github.com/gin-gonic/gin"
)

// 插入类型的业务逻辑
// 第一步接受用户的数据（Controller部分）
// 构造一个与数据一致的对象来接受数据 json.Unmarshall（Controller部分）
// 第二步将对象转化成一个SQL语句 (Service部分)
// 连接数据库获得*DB对象，然后拼接SQL
// 直接拼接SQL:fmt.SPrintf("INSERT INTO user VALUES(%s,%s,%s)",对象属性1，对象属性2，对象属性3)
// 通过db.Exec(SQL语句)完成插入数据的功能
// UserRequest 代表用户请求

// 发布一条消息 发布一条新闻 注册一个用户 都属于插入类型
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

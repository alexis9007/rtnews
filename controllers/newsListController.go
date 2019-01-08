package controllers

import (
	"fmt"
	"rtnews/services"

	"github.com/gin-gonic/gin"
)

// 查询类型的业务逻辑
// 如果是不带条件的查询，那么我们也没有必要接受用户的数据，直接查询数据库返回列表即可
// 通过业务层，连接数据库获得*DB,执行 'SELECT * FROM 某某表'
// 带条件的查询：
// 	条件是来自于用户请求数据的，一般情况下是在查询字符串里面
//  GET http://example.com?condition=value&condition2=value2... 多个条件通过&相连的，都在?后面
// 获得条件以后，要利用这些条件组装成一个完整SQL语句，利用
// rows,err := db.Query(SQL语句)来完成查询
// 查询的结果是包含所有行的一个结果，是一行一行的，但是程序是需要使用对象的。
// 所以才会需要 行数据 到 对象的转化，要利用 rows.Scan(属性名，属性名2，属性名3。。。)来完成转化
// 完成转化以后，将对象添加到slice结构里面，返回即可
func NewsListController(ctx *gin.Context) {
	authorVal, _ := ctx.GetQuery("author")
	data, err := services.NewsListService(fmt.Sprintf("author='%s'", authorVal))
	if err != nil {
		ctx.JSON(403, map[string]interface{}{
			"Code":    1001,
			"Message": "Failed",
			"Data":    nil,
		})
		return
	}
	ctx.JSON(200, map[string]interface{}{
		"Code":    1000,
		"Message": "Success",
		"Data":    data,
	})
	return
}

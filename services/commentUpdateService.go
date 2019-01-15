package services

import (
	"database/sql"
	"fmt"
	"log"
)

// TODO:
// 1、文件名统一为小写开头 RetriveCommentService -> retriveCommentService
// 2、Update是更新的意思 Upload才是上传 发布的意思
func CommentUpdate(content string, authorId int, newsId int) (bool, error) {
	db, err := sql.Open("mysql", "root:@/rtnews")
	defer db.Close()
	if err != nil {
		log.Print("sql open fatal error ", err)
		return false, err
	}
	//FIXME:
	//字符串里面占位符是有类型的 '%s' 代表传入的参数应该是个字符串，但是ID是整数类型的
	//应该更改为:
	/*
		stmt := fmt.Sprintf("INSERT INTO comment VALUES('%s','%d','%d')",conetn,authorId,newsId)
	*/
	stmt := fmt.Sprintf("INSERT INTO comment VALUES('%s','%s','%s')", content, authorId, newsId)
	log.Print("real sql statement is ", stmt)
	//FIXME:问题出现在这里 db.Exec(stmt)的返回类型是(Result,error)，没有对该处理进行处理，所以前端显示成功，但是没有插入成功
	//应该更改为
	/*
		_,err = db.Exec(stmt)
		if err != nil {
			return false,err
		}
		return true,nil
	*/
	db.Exec(stmt)
	return true, nil
}

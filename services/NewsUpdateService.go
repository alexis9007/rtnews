package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewsUpdateService(title string, content string, publishedAt time.Time, author string) (bool, error) {
	db, err := sql.Open("mysql", "root:@/rtnews")
	if err != nil {
		log.Println("call sqlOpen fatal error ", err)
		return false, err
	}
	stmt := fmt.Sprintf("INSER INTO news VALUES (3,'%s','%s','2019-01-10 16:29:27','%s')", title, content, author)
	log.Println("execute sql statement ", stmt)
	_, err = db.Exec(stmt)
	if err != nil {
		log.Printf("EXEC fatal error %v", err)
		return false, errors.New(fmt.Sprintf("插入数据出现严重错误%v", err))
	}
	return true, nil

}

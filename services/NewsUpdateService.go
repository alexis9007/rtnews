package services

import (
	"database/sql"
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
	stmt := fmt.Sprintf("INSER INTO news VALUES (0,%s, %s, %s, %s)", title, content, publishedAt, author)
	db.Exec(stmt)
	log.Println("execute sql statement ", stmt)
	return true, nil

}

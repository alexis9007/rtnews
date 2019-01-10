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
	stmt := fmt.Sprintf("INSER INTO news VALUES (0,%s, %s, %s, %s)", title, content, publishedAt.String(), author)
	log.Println("execute sql statement ", stmt)
	_, err = db.Exec(stmt)
	if err != nil {
		log.Printf("EXEC fatal error %v", err)
	}
	return true, nil

}

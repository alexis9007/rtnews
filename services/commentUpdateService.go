package services

import (
	"database/sql"
	"fmt"
	"log"
)

func CommentUpdate(content string, authorId int, newsId int) (bool, error) {
	db, err := sql.Open("mysql", "root:@/rtnews")
	defer db.Close()
	if err != nil {
		log.Print("sql open fatal error ", err)
		return false, err
	}

	stmt := fmt.Sprintf("INSERT INTO comment VALUES('%s','%d','%d')", content, authorId, newsId)
	log.Print("real sql statement is ", stmt)

	_, err = db.Exec(stmt)
	if err != nil {
		return false, err
	}
	return true, nil

}

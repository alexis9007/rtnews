package services

import (
	"database/sql"
	"fmt"
	"log"

	//"errors"
	_ "github.com/go-sql-driver/mysql"
)

func LoginService(username string, password string) (bool, error) {
	db, err := sql.Open("mysql", "root:@/rtnews")
	if err != nil {
		return false, err
	}

	stmt := fmt.Sprintf("SELECT username, password FROM user WHERE username='%s' AND password='%s'", username, password)
	log.Println("[INFO]real sql statement", stmt)
	rows, err := db.Query(stmt)
	if err != nil {
		return false, err
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil

}

package services

import (
	"database/sql"
	"fmt"
	"log"
	"rtnews/models"

	_ "github.com/go-sql-driver/mysql"
)

func UserListService(wheres string) ([]models.User, error) {
	db, err := sql.Open("mysql", "root:@/rtnews")
	defer db.Close()
	if err != nil {
		log.Println("sql open fatal error", err)
		return []models.User{}, err
	}

	stmt := fmt.Sprintf("SELECT username, password From user WHERE " + wheres)
	log.Println("execute sqlstatement is", stmt)
	row, err := db.Query(stmt)
	if err != nil {
		log.Println("Query from user fatal error", err)
		return []models.User{}, err
	}

	result := []models.User{}
	for row.Next() {
		user := models.User{}
		err := row.Scan(&user.Username, &user.Password)
		if err != nil {
			continue
		}
		result = append(result, user)
	}
	return result, nil
}

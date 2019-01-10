package services

import (
	"database/sql"
	"log"
	"rtnews/models"

	_ "github.com/go-sql-driver/mysql"
)

func NewsListService(wheres string) ([]models.News, error) {
	db, err := sql.Open("mysql", "root:@/rtnews")
	defer db.Close()
	if err != nil {
		return []models.News{}, err
	}
	//SELECT username,password from user
	stmt := "SELECT title,content,author from news WHERE " + wheres
	log.Println("executed sql statement is ", stmt)
	rows, err := db.Query(stmt)
	if err != nil {
		log.Printf("QUERY from news list fatal error %v", err)
		return []models.News{}, err
	}
	//result := []models.User{}
	result := []models.News{}
	//for rows.Next() {
	//	user := models.User{}
	//  err := rows.Scan(&user.Username,&user.Password)
	//  result = append(result,user)
	//}
	// return result
	for rows.Next() {
		news := models.News{}
		err := rows.Scan(&news.Title, &news.Content, &news.Author)
		if err != nil {
			continue
		}
		result = append(result, news)
	}
	return result, nil
}

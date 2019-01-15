package services

import (
	"database/sql"
	"log"
	"rtnews/models"
)

func RetriveCommentService(where string) ([]models.Comment, error) {
	db, err := sql.Open("mysql", "root:@/rtnews")
	defer db.Close()
	if err != nil {
		log.Print("sql open fatal error ", err)
		return []models.Comment{}, err
	}
	stmt := "SELECT content, athuorId, newsId FROM comment WHERE " + where
	log.Print("execute sql statement ", stmt)
	rows, err := db.Query(stmt)
	if err != nil {
		log.Print("Query from comment fatal error ", err)
		return []models.Comment{}, err
	}

	result := []models.Comment{}
	for rows.Next() {
		comment := models.Comment{}
		err := rows.Scan(&comment.Content, &comment.AuthorId, &comment.NewsId)
		if err != nil {
			continue
		}
		result = append(result, comment)
	}
	return result, nil
}

package services

import (
	"log"
	"fmt"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

// Reirements 1.0.0
func RegisterService(username string,password string) (bool,error) {
	//gorm => db
	// exist := db.Find("username=?",username)
	// if exist { return false }
	// if len(password) <= 6 {
	//	return false	
	//}
	// insertOk := db.Create(&User{username,password})
	// if !insertOk { return false }
	// return true
	db,err := sql.Open("mysql","root:@/rtnews")
	if err != nil {
		return false,err
	}
	stmt := fmt.Sprintf("INSERT INTO user VALUES('%s','%s')",username,password)
	log.Println("[INFO]real sql statement",stmt)
	db.Exec(stmt)
	return true,errors.New("password is too short")
}




package tests

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// func TestConnectMysql(t *testing.T) {
// 	db,err := sql.Open("mysql","root:@/rtnews")
// 	if err != nil {
// 		panic(err)
// 	}else{
// 		db.Exec("CREATE TABLE user(username varchar(255),password varchar(255)); ")
// 	}
// }

func insertIntoUser(un string, pw string) string {
	return fmt.Sprintf("INSERT INTO user VALUES('%s','%s')", un, pw)
}

func insertIntoAnyTable(tb, un, pw string) string {
	return fmt.Sprintf("INSERT INTO %s VALUES('%s','%s')", tb, un, pw)
}

func TestInsert(t *testing.T) {
	t.Log(time.Now().String())
	db, err := sql.Open("mysql", "root:@/rtnews")
	if err != nil {
		panic(err)
	}
	stmt := insertIntoAnyTable("user", "Li", "1234")
	db.Exec(stmt)
}

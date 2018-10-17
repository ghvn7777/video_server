package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


var (
	dbConn *sql.DB
	err error
)

func init() {
	fmt.Println("open and connect mysql")
	// 一般会做成配置文件
	dbConn, err = sql.Open("mysql", "root:ha@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	// defer dbConn.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = dbConn.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
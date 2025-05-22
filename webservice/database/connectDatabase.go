package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func SetUpDB() {
	var err error
	Db, err = sql.Open("mysql", "root:MayuSQL3310@tcp(127.0.0.1:3306)/stitchdb")
	if err != nil {
		fmt.Println("fail to connect")
	} else {
		fmt.Println("connect successfully")
	}
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxIdleConns(10)
	Db.SetConnMaxIdleTime(10)
}

package database

import (
	"database/sql"
	"fmt"
	"time"

	forenv "github.com/MayuraJam/StitchAlienAPIGO/forEnv"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func SetUpDB() {
	var err error
	dbUrl := forenv.EnvVatiable("DATABASE_URL")
	Db, err = sql.Open("mysql", dbUrl)

	if err != nil {
		fmt.Println("fail to connect")
	} else {
		fmt.Println("connect successfully")
	}
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxIdleConns(10)
	Db.SetConnMaxIdleTime(10)
}

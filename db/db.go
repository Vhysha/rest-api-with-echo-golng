package db

import (
	"database/sql"

	"github.com/AhmadShaadiqBahar/rest-api-with-echo-golng/config"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connection error")
	}

	err = db.Ping()
	if err != nil {
		panic("hhh")
	}

}

func CreateCon() *sql.DB {
	return db
}

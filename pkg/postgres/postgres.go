package postgres

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	// get data `host`, `user`, `password`, `database name`, and `port` from env
	var DB_HOST = os.Getenv("DB_HOST")
	var DB_USER = os.Getenv("DB_USER")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_PORT = os.Getenv("DB_PORT")

	// open connection to database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Jakarta", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// set max connections, cause gorm not hava default max connections value
	db, _ := DB.DB()
	db.SetMaxOpenConns(100)

	fmt.Println("Connected to Postgres Database")
}

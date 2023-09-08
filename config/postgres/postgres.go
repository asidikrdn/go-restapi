package postgres

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	// get data `host`, `user`, `password`, `database name`, and `port` from env
	var dbHost = os.Getenv("DB_HOST")
	var dbUser = os.Getenv("DB_USER")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbName = os.Getenv("DB_NAME")
	var dbPort = os.Getenv("DB_PORT")

	// open connection to database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Jakarta", dbHost, dbUser, dbPassword, dbName, dbPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// set max connections, cause gorm not hava default max connections value
	db, _ := DB.DB()
	db.SetMaxOpenConns(5)

	log.Println("Connected to Postgres Database")
}

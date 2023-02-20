package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database definition
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	// open database connection
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// set max concurrent connections to 5, if this app have more than 5 concurrent connections, several connections must wait till one of 5 concurrent connection finished its process.
	// If not set then this app have more than 5 concurrent connections, several connections may be error
	db, _ := DB.DB()
	db.SetMaxOpenConns(5)

	fmt.Println("Connected to MySQL database")
}

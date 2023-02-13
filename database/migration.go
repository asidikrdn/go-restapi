package database

import (
	"fmt"
	"go-restapi-boilerplate/pkg/postgre"
	"log"
)

func RunMigration() {
	// run auto migration
	err := postgre.DB.AutoMigrate(
	// put all models struct here
	// ex : &models.User{}
	)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Migration failed")
	}

	fmt.Println("Migration completed successfully")
}

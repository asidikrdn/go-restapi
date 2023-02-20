package database

import (
	"fmt"
	"go-restapi-boilerplate/pkg/postgres"
	"log"
)

func RunMigration() {
	// run auto migration
	err := postgres.DB.AutoMigrate(
	// put all models struct here
	// ex : &models.User{}
	)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Migration failed")
	}

	fmt.Println("Migration completed successfully")
}

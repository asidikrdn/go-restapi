package migrations

import (
	"fmt"
	"go-restapi-boilerplate/config/postgres"
	"go-restapi-boilerplate/db/models"
	"log"
)

// migration up
func RunMigration() {
	err := postgres.DB.AutoMigrate(
		&models.Log{},
		&models.MstRole{}, &models.MstUser{},
		// put another models struct here
	)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Migration failed")
	}

	fmt.Println("Migration up completed successfully")
}

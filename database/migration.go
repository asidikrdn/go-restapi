package database

import (
	"fmt"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/pkg/postgres"
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

// migration down
func DropMigration() {
	err := postgres.DB.Migrator().DropTable(
		&models.Log{},
		&models.MstRole{}, &models.MstUser{},
		// put another models struct here
	)

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Migration failed")
	}

	fmt.Println("Migration down completed successfully")
}

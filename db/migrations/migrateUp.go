package migrations

import (
	"go-restapi/config/postgres"
	"go-restapi/db/models"
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
		log.Println(err.Error())
		log.Fatal("Migration failed")
	}

	log.Println("Migration up completed successfully")
}

package migrations

import (
	"go-restapi/config/postgres"
	"go-restapi/db/models"
	"log"
)

// migration down
func DropMigration() {
	err := postgres.DB.Migrator().DropTable(
		&models.Log{},
		&models.MstRole{}, &models.MstUser{},
		// put another models struct here
	)

	if err != nil {
		log.Println(err.Error())
		log.Fatal("Migration failed")
	}

	log.Println("Migration down completed successfully")
}

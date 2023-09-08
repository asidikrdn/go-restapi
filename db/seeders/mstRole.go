package seeders

import (
	"go-restapi/config/postgres"
	"go-restapi/db/models"
	"log"
)

func seedMstRole() {
	// add some role
	if postgres.DB.Migrator().HasTable(&models.MstRole{}) {
		newRole := []models.MstRole{}

		newRole = append(newRole, models.MstRole{
			Role: "Superadmin",
		})
		newRole = append(newRole, models.MstRole{
			Role: "Admin",
		})
		newRole = append(newRole, models.MstRole{
			Role: "User",
		})

		for _, role := range newRole {
			errAddRole := postgres.DB.Create(&role).Error
			if errAddRole != nil {
				log.Println(errAddRole.Error())
				log.Fatal("Seeding failed")
			}
		}

		log.Println("Success seeding master role...")
	}
}

package seeders

import (
	"fmt"
	"go-restapi-boilerplate/config/postgres"
	"go-restapi-boilerplate/db/models"
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
				fmt.Println(errAddRole.Error())
				log.Fatal("Seeding failed")
			}
		}

		fmt.Println("Success seeding master role...")
	}
}

package database

import (
	"fmt"
)

func RunSeeder() {
	// ==================================
	// EXAMPLE
	// ==================================

	/*
		// cek is user table exist
		if postgres.DB.Migrator().HasTable(&models.User{}) {
			// check is user table has minimum 1 user
			err := postgres.DB.First(&models.User{}).Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// create 1 user
				newUser := models.User{
					FullName:        "Sidik",
					Email:           "sidik@mail.com",
				}

				hashPassword, err := bcrypt.HashingPassword("12345678")
				if err != nil {
					log.Fatal("Hash password failed")
				}

				newUser.Password = hashPassword

				// insert user to database
				errAddUser := postgres.DB.Create(&newUser).Error
				if errAddUser != nil {
					fmt.Println(errAddUser.Error())
					log.Fatal("Seeding failed")
				}
			}
		}
	*/

	fmt.Println("Seeding completed successfully")
}

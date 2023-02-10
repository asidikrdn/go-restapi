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
		if mysql.DB.Migrator().HasTable(&models.User{}) {
			// check is user table has minimum 1 user as admin
			err := mysql.DB.First(&models.User{}, "role = ?", "superadmin").Error
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// create 1 admin
				newUser := models.User{
					UserID:          "000001",
					FullName:        "Administrator",
					Role:            "superadmin",
					Email:           "admin@sigesit.com",
					IsEmailVerified: false,
				}

				hashPassword, err := bcrypt.HashingPassword("12345678")
				if err != nil {
					log.Fatal("Hash password failed")
				}

				newUser.Password = hashPassword

				// insert admin to database
				errAddUser := mysql.DB.Select("UserID", "FullName", "Role", "Email", "IsEmailVerified", "Password").Create(&newUser).Error
				if errAddUser != nil {
					fmt.Println(errAddUser.Error())
					log.Fatal("Seeding failed")
				}
			}
		}
	*/

	fmt.Println("Seeding completed successfully")
}

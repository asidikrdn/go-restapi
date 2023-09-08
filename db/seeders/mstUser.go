package seeders

import (
	"errors"
	"fmt"
	"go-restapi/config/postgres"
	"go-restapi/db/models"
	"go-restapi/pkg/bcrypt"
	"log"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func seedMstUser() {
	if postgres.DB.Migrator().HasTable(&models.MstUser{}) {
		// check is user table has minimum 1 user
		err := postgres.DB.First(&models.MstUser{}).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// add superadmin
			newUser := models.MstUser{
				ID:              uuid.New(),
				FullName:        "Super Admin",
				Email:           os.Getenv("SUPERADMIN_EMAIL"),
				IsEmailVerified: true,
				RoleID:          1,
			}

			hashPassword, err := bcrypt.HashingPassword(os.Getenv("SUPERADMIN_PASSWORD"))
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
		fmt.Println("Success seeding super admin...")
	}
}

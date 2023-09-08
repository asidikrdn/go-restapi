package seeders

import "log"

func RunSeeder() {
	// Role
	seedMstRole()

	// User
	seedMstUser()

	log.Println("Seeding completed successfully")
}

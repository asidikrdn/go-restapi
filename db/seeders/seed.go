package seeders

import (
	"fmt"
)

func RunSeeder() {
	// Role
	seedMstRole()

	// User
	seedMstUser()

	fmt.Println("Seeding completed successfully")
}

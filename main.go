package main

import (
	"fmt"
	"go-restapi-boilerplate/database"
	"go-restapi-boilerplate/pkg/postgres"
	"go-restapi-boilerplate/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environment variables file, the apps will read global environtment variabels on this system")
	}

	// database initialization
	postgres.DatabaseInit()
	// redis.RedisInit()

	// database migration & seeder
	database.DropMigration()
	database.RunMigration()
	database.RunSeeder()

	// gin Mode
	gin.SetMode(os.Getenv("GIN_MODE"))

	// create new router
	router := gin.Default()

	// call routerinit with pathprefix
	routes.RouterInit(router.Group("/api/v1"))

	// file server endpoint
	router.Static("/static", "./uploads")

	//	set up CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://example.com"} // Replace with your allowed origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}

	// Add cors middleware on all route
	router.Use(cors.New(config))

	// Running services
	fmt.Println("server running on localhost:" + os.Getenv("PORT"))
	router.Run(":" + os.Getenv("PORT"))
}

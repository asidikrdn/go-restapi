package main

import (
	"fmt"
	"go-restapi-boilerplate/database"
	"go-restapi-boilerplate/pkg/middleware"
	"go-restapi-boilerplate/pkg/postgres"
	"go-restapi-boilerplate/routes"
	"os"
	"strconv"

	"github.com/asidikrdn/otptimize"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load environment variables
	err := godotenv.Load("config.env")
	if err != nil {
		fmt.Println("Error loading environment variables file, the apps will read global environtment variabels on this system")
	}

	// database initialization
	postgres.DatabaseInit()

	// otptimize connection init
	mailPort, _ := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	mailConfig := otptimize.MailConfig{
		Host:     os.Getenv("CONFIG_SMTP_HOST"),
		Port:     mailPort,
		Email:    os.Getenv("CONFIG_AUTH_EMAIL"),
		Password: os.Getenv("CONFIG_AUTH_PASSWORD"),
	}
	redisConfig := otptimize.RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
	otptimize.ConnectionInit(mailConfig, redisConfig)

	// database migration & seeder
	database.DropMigration()
	database.RunMigration()
	database.RunSeeder()

	// gin Mode
	gin.SetMode(os.Getenv("GIN_MODE"))

	// create new router
	router := gin.Default()

	// call logger middleware before route to any routes
	router.Use(middleware.Logger())

	//	set up CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Replace with your allowed origins
	config.AllowMethods = []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Authorization"}

	// Add cors middleware on all route
	router.Use(cors.New(config))

	// call routerinit with pathprefix
	routes.RouterInit(router.Group("/api/v1"))

	// file server endpoint
	router.Static("/static", "./uploads")

	// Running services
	fmt.Println("server running on localhost:" + os.Getenv("PORT"))
	router.Run(":" + os.Getenv("PORT"))
}

package main

import (
	"fmt"
	"go-restapi-boilerplate/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	// load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environment variables file, the apps will read global environtment variabels on this system")
	}

	// database initialization
	// mysql.DatabaseInit()
	// redis.RedisInit()

	// database migration
	// database.RunMigration()

	// gin Mode
	gin.SetMode(os.Getenv("GIN_MODE"))

	// create new router
	router := gin.Default()

	// call routerinit with pathprefix
	routes.RouterInit(router.Group("/api/v1"))

	// file server endpoint
	router.Static("/static", "./uploads")

	//	set up CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedMethods:   []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowCredentials: true,
	})

	// Add cors middleware on all route
	handler := c.Handler(router)

	// Running services
	fmt.Println("server running on localhost:" + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), handler)
}

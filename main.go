package main

import (
	"fmt"
	"go-restapi-boilerplate/database"
	"go-restapi-boilerplate/pkg/mysql"
	"go-restapi-boilerplate/routes"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environment variables file, the apps will read global environtment variabels on this system")
	}

	// database initialization
	mysql.DatabaseInit()
	// redis.RedisInit()

	// database migration & seeder
	database.RunMigration()
	database.RunSeeder()

	// create new router
	router := mux.NewRouter()

	// call routerinit with pathprefix
	routes.RouterInit(router.PathPrefix("/api/v1").Subrouter())

	// file server endpoint
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./uploads"))))

	//	set up and add cors middleware on all route
	AllowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	AllowedOrigins := handlers.AllowedOrigins([]string{"*"})
	AllowedMethods := handlers.AllowedMethods([]string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"})

	// Running services
	fmt.Println("server running on localhost:" + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(AllowedHeaders, AllowedOrigins, AllowedMethods)(router))
}

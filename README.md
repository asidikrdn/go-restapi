# Go REST API Boilerplate

This boilerplate using uncle-bob clean architecture models and includes some basic third party package that will be used on any rest api.
If you want to use this boilerplate to develop your own project, don't forget to change the project name in `go.mod` file.

## Clean Architecture

The application follows the Uncle Bob "[Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)" principles and project structure :

### Clean Architecture layers

![Schema of flow of Clean Architecture](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

### Project anatomy

```text
app
  └ database                        → Migration & Seeder functions from models to database
  └ dto                             → Converter objects that transform outside objects (ex: HTTP request payload) to inside objects (ex: database models), or vice versa 
  └ handlers                        → Contains any function to handle the request and serve the response
  └ models                          → Initalize models that want to be migrated to database
  └ pkg                             → Frameworks, drivers and tools such as Database, custom middleware, mailing/logging/glue code etc
  └ repositories                    → Contains any function to communicate with database, such as get data or send data
  └ routes                          → Capture and redirect the incoming request to the specified handlers
  └ uploads                         → Contains any uploaded files/folder
  └ .dockerignore                   → Configuration file that describes files and directories that want to exclude when building a docker image
  └ .gitignore                      → Configuration file that describes files and directories that want to ignore when committing your project to the git repository
  └ Dockerfile                      → Docker setup configuration for building docker image
  └ go.mod                          → Contains project name and all the module dependencies which are needed or to be used in the project
  └ go.sum                          → Contains information used by Go to record specific hashes and versions of dependencies
  └ main.go                         → Main application entry point
```

routers → handlers → repository → database

## Tech Stack

Stack in this boilerplate :

- `gin-gonic` as main library to handle any routes.
- `bcrypt` as password encryption to secure the password before store it in the database.
- `jwt` as a token to authenticate communication from other service, such as client (front-end).
- `gorm` as ORM library to handle communication between the apps and SQL databases. And postgres driver to connect to postgreSQL.
- `godotenv` as an .env file reader to read the environment variables form .env file.
- `cors` as cors configuration.

## Environtment Variables

Environment Variables needed in this application :

```env
# database setup
DB_HOST=<your_database_host>
DB_PORT=<your_database_port>
DB_USER=<your_database_user>
DB_PASSWORD=<your_database_password>
DB_NAME=<your_database_name>

# redis setup
REDIS_HOST=<your_database_host>
REDIS_PORT=<your_database_port>
REDIS_PASSWORD=<your_database_password>

# gin mode
GIN_MODE=<release/debug>

# port for running apps
PORT=<your_port>
```

## SetUp

Before run this apps, you must install and run the PostgreSQL database. After that, create a new database with name depend on your .env file.
And this is step-by-step to run this application :

- Run on terminal `git clone https://github.com/asidikrdn/go-restapi-boilerplate.git`
- Run on terminal `go mod download`
- Edit and customize this app depending on your needs
- Run on terminal `go run main.go`

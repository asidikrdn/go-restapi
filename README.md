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
- `joho/godotenv` as an .env file reader to read the environment variables form .env file.
- `gin-contrib/cors` as cors configuration.

## Environtment Variables

The required environment variables to run this application are available in the .env.example file. Please copy that file to a new .env file and adjust its contents according to your needs

## API Documentation

- [View Postman Documentation](https://documenter.getpostman.com/view/23687279/2s9YBxXaR2)

<!-- ![Screenshot Dokumentasi Postman](postman/screenshot.png) -->

IThis is our API documentation created with Postman. This documentation provides a comprehensive description of all endpoints and how to use them. Make sure to refer to this documentation while developing or integrating with our API.

**Documentation Updates:** This documentation is periodically updated to reflect changes to our API. Be sure to always refer to the latest documentation.

## SetUp

Before run this apps, you must install and run the PostgreSQL database. After that, create a new database with name depend on your .env file.
And this is step-by-step to run this application :

1. Clone the repository by running the following command in your terminal:

   ```sh
   git clone https://github.com/asidikrdn/go-restapi-boilerplate.git
   ```

2. Navigate to the project directory:

   ```sh
   cd go-restapi-boilerplate
   ```

3. Install dependencies using Go Modules:

   ```sh
   go mod download
   ```

4. Duplicate the .env.example file and rename it to .env. Customize the values of the environment variables as needed.

5. Launch the application with the following command:

   ```sh
   go run main.go
   ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

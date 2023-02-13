# Go REST API Boilerplate

This boilerplate using uncle-bob clean architecture models and includes some basic third party package that will be used on any rest api.

## Tech Stack

Stack in this boilerplate :

- `gin-gonic` as main library to handle any requests and serve the response.
- `bcrypt` as password encryption to secure the password before store it in the database.
- `jwt` as a token to authenticate communication from other service, such as client (front-end).
- `gorm` as ORM library to handle communication between the apps and SQL databases. And postgree driver to connect to postgreSQL.
- `godotenv` as an .env file reader to read the environment variables form .env file.
- `cors` as cors configuration.

## Environtment Variables

Environment Variables needed in this application :

```s
# database setup
DB_HOST=<your_database_host>
DB_PORT=<your_database_port>
DB_USER=<your_database_user>
DB_PASSWORD=<your_database_password>
DB_NAME=<your_database_name>

#redis setup
REDIS_HOST=<your_database_host>
REDIS_PORT=<your_database_port>
REDIS_PASSWORD=<your_database_password>
```

If you want to use this boilerplate to develop your own project, don't forget to change the project name in `go.mod` file.

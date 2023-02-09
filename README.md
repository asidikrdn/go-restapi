# Go REST API Boilerplate

This boilerplate using uncle-bob clean architecture models and includes some basic package that will be used on any rest api, such as bcrypt and json web token.

Stack used in this boilerplate :

- `gin-gonic` as main library to handle any requests and serve the response.
- `bcrypt` as password encryption to secure the password before store it in the database.
- `jwt` as a token to authenticate communication from other service, such as client (front-end).
- `gorm` as ORM library to handle communication between the apps and SQL databases.
- `godotenv` as an .env file reader to read the environment variables form .env file.
- `cors` as cors configuration.

If you want to use this boilerplate to develop your own project, don't forget to change the project name in `go.mod` file.

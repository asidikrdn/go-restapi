# ROUTES

This folder contains several route that separated into several files. Any routes on each file will be combined into a single file named `router.go`. `router.go` will be called in main application and passing main router object to any routes. Main function of router is to capture and redirect the incoming request to the specified handlers.

example :
`user.go`

```go
package routes

import (
 "go-restapi-boilerplate/handlers"
 "go-restapi-boilerplate/pkg/middleware"
 "go-restapi-boilerplate/pkg/mysql"
 "go-restapi-boilerplate/repositories"

 "github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup) {
 userRepository := repositories.MakeRepository(mysql.DB)
 h := handlers.HandlerUser(userRepository)

 //  without middleware
 r.GET("/users", h.FindAllCustomer)
 //  with middleware
 r.GET("/users", middleware.AdminAuth(), h.FindAllCustomer)
 }
```

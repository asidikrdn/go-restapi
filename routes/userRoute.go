package routes

import (
	"go-restapi-boilerplate/handlers/handlerUser"
	"go-restapi-boilerplate/pkg/middleware"
	"go-restapi-boilerplate/pkg/postgres"
	"go-restapi-boilerplate/repositories"

	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup) {
	userRepository := repositories.MakeRepository(postgres.DB)
	h := handlerUser.HandlerUser(userRepository)

	// login
	r.POST("/login", h.Login)

	// check auth
	r.GET("/check-auth", middleware.UserAuth(), h.CheckAuth)

	// create new user
	r.POST("/users", middleware.AdminAuth(), middleware.UploadSingleFile(), h.CreateUser)

	// find/get user
	r.GET("/users", middleware.AdminAuth(), h.FindAllUsers)
	r.GET("/users/:id", middleware.AdminAuth(), h.FindUserByID)
	r.GET("/users/profile", middleware.UserAuth(), h.GetProfile)

	// update user
	r.PATCH("/users/:id", middleware.AdminAuth(), middleware.UploadSingleFile(), h.UpdateUserByID)
	r.PATCH("/users/profile", middleware.UserAuth(), middleware.UploadSingleFile(), h.UpdateProfile)

	// delete user
	r.DELETE("/users/:id", middleware.AdminAuth(), h.DeleteUser)
}

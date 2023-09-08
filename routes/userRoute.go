package routes

import (
	"go-restapi/config/postgres"
	"go-restapi/handlers/handlerUser"
	"go-restapi/pkg/middleware"
	"go-restapi/repositories"

	"github.com/gin-gonic/gin"
)

var userByID = "/users/:id"

func User(r *gin.RouterGroup) {
	userRepository := repositories.MakeRepository(postgres.DB)
	h := handlerUser.HandlerUser(userRepository)

	// create user by admin/superadmin
	r.POST("/users", middleware.AdminAuth(), middleware.UploadSingleFile(), h.CreateUserByAdmin)

	// find/get user
	r.GET("/users", middleware.AdminAuth(), h.FindAllUsers)
	r.GET(userByID, middleware.AdminAuth(), h.FindUserByID)
	r.GET("/users/profile", middleware.UserAuth(), h.GetProfile)

	// update user
	r.PATCH(userByID, middleware.AdminAuth(), middleware.UploadSingleFile(), h.UpdateUserByID)
	r.PATCH("/users/profile", middleware.UserAuth(), middleware.UploadSingleFile(), h.UpdateProfile)

	// delete user
	r.DELETE(userByID, middleware.AdminAuth(), h.DeleteUser)
}

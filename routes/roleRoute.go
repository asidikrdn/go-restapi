package routes

import (
	"go-restapi-boilerplate/handlers/handlerRole"
	"go-restapi-boilerplate/pkg/middleware"
	"go-restapi-boilerplate/pkg/postgres"
	"go-restapi-boilerplate/repositories"

	"github.com/gin-gonic/gin"
)

func Role(r *gin.RouterGroup) {
	roleRepository := repositories.MakeRepository(postgres.DB)
	h := handlerRole.HandlerRole(roleRepository)

	// find all role
	r.GET("/roles", middleware.UserAuth(), h.FindAllRole)

	// find role by id
	r.GET("/roles/:id", middleware.UserAuth(), h.FindRoleByID)
}

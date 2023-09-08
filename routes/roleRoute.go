package routes

import (
	"go-restapi/config/postgres"
	"go-restapi/handlers/handlerRole"
	"go-restapi/pkg/middleware"
	"go-restapi/repositories"

	"github.com/gin-gonic/gin"
)

func Role(r *gin.RouterGroup) {
	roleRepository := repositories.MakeRepository(postgres.DB)
	h := handlerRole.HandlerRole(roleRepository)

	roleByID := "/roles/:id"

	// find all role
	r.GET("/roles", middleware.AdminAuth(), h.FindAllRole)

	// find role by id
	r.GET(roleByID, middleware.AdminAuth(), h.FindRoleByID)

	// add new role
	r.POST("/roles", middleware.SuperAdminAuth(), h.CreateRoleBySuperadmin)

	// update role
	r.PATCH(roleByID, middleware.SuperAdminAuth(), h.UpdateRoleByID)

	// delete role
	r.DELETE(roleByID, middleware.SuperAdminAuth(), h.DeleteRoleByID)
}

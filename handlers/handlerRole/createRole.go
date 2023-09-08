package handlerRole

import (
	"go-restapi-boilerplate/db/models"
	"go-restapi-boilerplate/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handlerRole) CreateRoleBySuperadmin(c *gin.Context) {
	var request dto.CreateRoleRequest

	// get request data
	err := c.ShouldBind(&request)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create new role
	role := models.MstRole{
		Role: request.Role,
	}

	// save new role data to database
	addedRole, err := h.RoleRepository.CreateRole(&role)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// reload data
	newRole, err := h.RoleRepository.FindRoleByID(addedRole.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// send response
	response := dto.Result{
		Status:  http.StatusCreated,
		Message: "OK",
		Data:    convertRoleResponse(newRole),
	}
	c.JSON(http.StatusCreated, response)
}

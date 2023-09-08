package handlerRole

import (
	"go-restapi/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlerRole) UpdateRoleByID(c *gin.Context) {
	var request dto.UpdateRoleRequest

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

	// get id from url
	id, _ := strconv.Atoi(c.Param("id"))

	// get role data from database
	role, err := h.RoleRepository.FindRoleByID(uint(id))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// update role
	if request.Role != "" && request.Role != role.Role {
		role.Role = request.Role
	}

	// save updated role to database
	updatedRole, err := h.RoleRepository.UpdateRole(role)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// reload data
	role, err = h.RoleRepository.FindRoleByID(updatedRole.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertRoleResponse(role),
	}
	c.JSON(http.StatusOK, response)
}

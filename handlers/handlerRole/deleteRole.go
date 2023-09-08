package handlerRole

import (
	"go-restapi/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlerRole) DeleteRoleByID(c *gin.Context) {
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

	// check is role used by user
	roleUsed, err := h.RoleRepository.CheckIsRoleUsed(role)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	} else if roleUsed {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "can't delete role, cause role is used by some user",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// delete role from database
	role, err = h.RoleRepository.DeleteRole(role)
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
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertRoleResponse(role),
	}
	c.JSON(http.StatusOK, response)
}

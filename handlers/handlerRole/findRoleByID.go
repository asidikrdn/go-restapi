package handlerRole

import (
	"go-restapi/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlerRole) FindRoleByID(c *gin.Context) {
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

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertRoleResponse(role),
	}
	c.JSON(http.StatusOK, response)
}

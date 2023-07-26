package handlerUser

import (
	"fmt"
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handlerUser) DeleteUser(c *gin.Context) {
	// get userid from url
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// get user data
	user, err := h.UserRepository.FindUserByID(id)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// delete user image if exist
	if user.Image != "" {
		if !helpers.DeleteFile(user.Image) {
			fmt.Println(err.Error())
		}
	}

	// delete user data
	user, err = h.UserRepository.DeleteUser(user)
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
		Data:    convertUserResponse(user),
	}
	c.JSON(http.StatusOK, response)
}

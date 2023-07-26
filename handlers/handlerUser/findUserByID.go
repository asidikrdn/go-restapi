package handlerUser

import (
	"go-restapi-boilerplate/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (h *handlerUser) FindUserByID(c *gin.Context) {
	// get userid
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

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertUserResponse(user),
	}
	c.JSON(http.StatusOK, response)
}

func (h *handlerUser) GetProfile(c *gin.Context) {
	// get jwt payload
	claims, ok := c.Get("userData")
	if !ok {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User data not found",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// extract user data from jwt claims
	userData := claims.(jwt.MapClaims)

	// get userid
	id, err := uuid.Parse(userData["id"].(string))
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

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertUserResponse(user),
	}
	c.JSON(http.StatusOK, response)
}

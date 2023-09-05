package handlerAuth

import (
	"go-restapi-boilerplate/dto"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (h *handlerAuth) CheckAuth(c *gin.Context) {
	// get jwt payload
	claims, ok := c.Get("userData")
	if !ok {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User data from jwt payload is not found",
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

	// find user data
	user, err := h.UserRepository.FindUserByID(id)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User not found",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertLoginResponse(user, strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", -1)),
	}
	c.JSON(http.StatusOK, response)
}

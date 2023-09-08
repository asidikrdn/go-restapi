package handlerAuth

import (
	"go-restapi/dto"
	"go-restapi/pkg/bcrypt"
	jwtToken "go-restapi/pkg/jwt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (h *handlerAuth) Login(c *gin.Context) {
	var request dto.LoginRequest

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

	// find user data
	user, err := h.UserRepository.GetUserByEmailOrPhone(request.Username)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User not found",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// check password
	if isPasswordValid := bcrypt.CheckPassword(request.Password, user.Password); !isPasswordValid {
		response := dto.Result{
			Status:  http.StatusUnauthorized,
			Message: "Password invalid",
		}
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	// preparing jwt claims
	myClaims := jwt.MapClaims{}
	myClaims["id"] = user.ID
	myClaims["fullname"] = user.FullName
	myClaims["email"] = user.Email
	myClaims["roleId"] = user.RoleID
	myClaims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24 hours expired

	// generate token
	token, err := jwtToken.GenerateToken(&myClaims)
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
		Data:    convertLoginResponse(user, token),
	}
	c.JSON(http.StatusOK, response)
}

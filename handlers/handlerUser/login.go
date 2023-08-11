package handlerUser

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/pkg/bcrypt"
	jwtToken "go-restapi-boilerplate/pkg/jwt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (h *handlerUser) Login(c *gin.Context) {
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
	myClaims["name"] = user.FullName
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
		Data: dto.LoginResponse{
			ID:    user.ID,
			Email: user.Email,
			Role: dto.RoleResponse{
				ID:   user.Role.ID,
				Role: user.Role.Role,
			},
			Token: token,
		},
	}
	c.JSON(http.StatusOK, response)
}

func (h *handlerUser) CheckAuth(c *gin.Context) {
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

	// preparing jwt claims
	myClaims := jwt.MapClaims{}
	myClaims["id"] = user.ID
	myClaims["name"] = user.FullName
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
		Data: dto.LoginResponse{
			ID:    user.ID,
			Email: user.Email,
			Role: dto.RoleResponse{
				ID:   user.Role.ID,
				Role: user.Role.Role,
			},
			Token: token,
		},
	}
	c.JSON(http.StatusOK, response)
}

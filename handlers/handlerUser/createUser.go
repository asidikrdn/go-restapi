package handlerUser

import (
	"go-restapi/db/models"
	"go-restapi/dto"
	"go-restapi/pkg/bcrypt"
	"net/http"
	"os"

	"github.com/asidikrdn/otptimize"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (h *handlerUser) CreateUserByAdmin(c *gin.Context) {
	var request dto.RegisterRequest

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

	// read user data from jwt payload
	userData := claims.(jwt.MapClaims)
	id, err := uuid.Parse(userData["id"].(string))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// get admin data from database
	admin, err := h.UserRepository.FindUserByID(id)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// check email
	_, err = h.UserRepository.GetUserByEmailOrPhone(request.Email)
	if err == nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email already registered",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// check phone
	_, err = h.UserRepository.GetUserByEmailOrPhone(request.Phone)
	if err == nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Phone number already registered",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create new user
	user := models.MstUser{
		ID:              uuid.New(),
		FullName:        request.FullName,
		Email:           request.Email,
		IsEmailVerified: false,
		Phone:           request.Phone,
		IsPhoneVerified: false,
		Address:         request.Address,
		// RoleID:          request.RoleID,
	}

	// admin only can create user, but superadmin can create user & admin
	createRole(admin, &user, request.RoleID)

	// hashing password
	user.Password, err = bcrypt.HashingPassword(request.Password)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// get image from context
	image, ok := c.Get("image")
	if ok {
		user.Image = image.(string)
	}

	// save new user data to database
	addedUser, err := h.UserRepository.CreateUser(&user)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// generate and send otp
	go otptimize.GenerateAndSendOTP(4, 5, os.Getenv("APP_NAME"), user.FullName, user.Email)

	// reload data
	newUser, err := h.UserRepository.FindUserByID(addedUser.ID)
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
		Data:    convertUserResponse(newUser),
	}
	c.JSON(http.StatusCreated, response)
}

func createRole(admin *models.MstUser, user *models.MstUser, requestData uint) {
	if admin.RoleID != 1 {
		user.RoleID = 3
	} else if admin.RoleID == 1 {
		user.RoleID = requestData
	}
}

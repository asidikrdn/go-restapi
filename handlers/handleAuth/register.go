package handlerAuth

import (
	"go-restapi-boilerplate/db/models"
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/pkg/bcrypt"
	"net/http"
	"os"

	"github.com/asidikrdn/otptimize"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handlerAuth) RegisterUser(c *gin.Context) {
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
		RoleID:          3,
	}

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
		Data:    convertRegisterResponse(newUser),
	}
	c.JSON(http.StatusCreated, response)
}

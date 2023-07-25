package handlerUser

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/pkg/bcrypt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handlerUser) CreateUser(c *gin.Context) {
	var request dto.CreateUserRequest

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

	user := models.MstUser{
		ID:              uuid.New(),
		FullName:        request.FullName,
		Email:           request.Email,
		IsEmailVerified: false,
		Phone:           request.Phone,
		IsPhoneVerified: false,
		Address:         request.Address,
		RoleID:          request.RoleID,
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

	addedUser, err := h.UserRepository.CreateUser(&user)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// generate OTP Code
	// otpCode := helpers.GenerateRandomOTP(4)

	// Send OTP Code to user's email
	// go helpers.SendVerificationEmail(otpCode, &user)

	// save hashed OTP Code on redis server
	// hashedOTP, err := bcrypt.HashingPassword(otpCode)
	// if err == nil {
	// 	// for email verification
	// 	err = redis.SetValue(user.Email, hashedOTP, time.Minute*5)
	// 	if err != nil {
	// 		fmt.Println("Failed to set value")
	// 	}
	// 	// for phone verification
	// 	err = redis.SetValue(user.Phone, hashedOTP, time.Minute*5)
	// 	if err != nil {
	// 		fmt.Println("Failed to set value")
	// 	}
	// }

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

	response := dto.Result{
		Status:  http.StatusCreated,
		Message: "OK",
		Data:    convertUserResponse(newUser),
	}
	c.JSON(http.StatusCreated, response)
}

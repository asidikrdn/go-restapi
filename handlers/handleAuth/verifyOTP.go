package handlerAuth

import (
	"go-restapi/dto"
	"net/http"

	"github.com/asidikrdn/otptimize"
	"github.com/gin-gonic/gin"
)

func (h *handlerAuth) VerifyEmail(c *gin.Context) {
	// get token from request
	var request dto.VerifyEmailRequest
	err := c.ShouldBind(&request)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get user data
	user, err := h.UserRepository.GetUserByEmailOrPhone(request.Email)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "User not found",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// verify otp
	isOtpValid, err := otptimize.ValidateOTP(request.Email, request.OTPToken)
	if !isOtpValid {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "OTP invalid",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	} else if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// change verification status
	user.IsEmailVerified = true

	// update verification status on database
	_, err = h.UserRepository.UpdateUser(user)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "User's email verified",
		Data:    convertRegisterResponse(user),
	}
	c.JSON(http.StatusOK, response)
}

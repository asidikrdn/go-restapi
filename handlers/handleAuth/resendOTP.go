package handlerAuth

import (
	"go-restapi/dto"
	"net/http"
	"os"
	"regexp"

	"github.com/asidikrdn/otptimize"
	"github.com/gin-gonic/gin"
)

func (h *handlerAuth) ResendOTP(c *gin.Context) {
	// get request data
	email := c.Param("email")

	// Ekspresi reguler untuk validasi email
	regexStr := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile regex
	regex := regexp.MustCompile(regexStr)

	// check email
	if !regex.MatchString(email) {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email invalid",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// check user's email, is it already registered
	user, err := h.UserRepository.GetUserByEmailOrPhone(email)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email not registered, please register first",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// email already verified
	if user.IsEmailVerified {
		response := dto.Result{
			Status:  http.StatusBadRequest,
			Message: "Email already verified",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// send new otp
	otptimize.GenerateAndSendOTP(4, 5, os.Getenv("APP_NAME"), user.FullName, user.Email)

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OTP has been sent successfully",
	}
	c.JSON(http.StatusOK, response)
}

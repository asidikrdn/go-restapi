package routes

import (
	"go-restapi/config/postgres"
	handlerAuth "go-restapi/handlers/handleAuth"
	"go-restapi/pkg/middleware"
	"go-restapi/repositories"

	"github.com/gin-gonic/gin"
)

func Auth(r *gin.RouterGroup) {

	userRepository := repositories.MakeRepository(postgres.DB)
	h := handlerAuth.HandlerAuth(userRepository)

	// login
	r.POST("/login", h.Login)

	// check auth
	r.GET("/check-auth", middleware.UserAuth(), h.CheckAuth)

	// register new user
	r.POST("/register", middleware.UploadSingleFile(), h.RegisterUser)

	// resend OTP
	r.GET("/otp/resend/:email", h.ResendOTP)

	// verify OTP
	r.POST("/otp/verify", h.VerifyEmail)
}

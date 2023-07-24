package middleware

import (
	"go-restapi-boilerplate/dto"
	jwtToken "go-restapi-boilerplate/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// mengambil token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort() // prevent doing next handler
			return
		}

		// validasi token dan mengambil claims
		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort() // prevent doing next handler
			return
		}

		// set up context value and send it to next handler
		c.Set("userData", claims)
		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort() // prevent doing next handler
			return
		}

		// validate token and get claims
		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort() // prevent doing next handler
			return
		}

		// validate is it admin
		if claims["roleId"].(float64) != 1 && claims["roleId"].(float64) != 2 {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized, you're not administrator",
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort() // prevent doing next handler
			return
		}

		// set up context value and send it to next handler
		c.Set("userData", claims)
		c.Next()
	}
}

func SuperAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort() // prevent doing next handler
			return
		}

		// validate token and get claims
		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort() // prevent doing next handler
			return
		}

		// validate is it superadmin
		if claims["roleId"].(float64) != 1 {
			response := dto.Result{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized, you're not Super Administrator",
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort() // prevent doing next handler
			return
		}

		// set up context value and send it to next handler
		c.Set("userData", claims)
		c.Next()
	}
}

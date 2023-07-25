package handlerUser

import (
	"fmt"
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (h *handlerUser) UpdateUserByID(c *gin.Context) {
	var request dto.UpdateUserRequest

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

	// get user id
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// get user data from database
	user, err := h.UserRepository.FindUserByID(id)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// update fullname
	if request.FullName != "" {
		user.FullName = request.FullName
	}

	// update email
	if request.Email != "" {
		user.Email = request.Email
		user.IsEmailVerified = false
	}

	// update phone
	if request.Phone != "" {
		user.Phone = request.Phone
		user.IsPhoneVerified = false
	}

	// update address
	if request.Address != "" {
		user.Address = request.Address
	}

	// update role
	if request.RoleID != 0 {
		user.RoleID = request.RoleID
	}

	// update image
	image, ok := c.Get("image")
	if ok {
		if user.Image != "" {
			if !helpers.DeleteFile(user.Image) {
				fmt.Println(err.Error())
			}
		}

		user.Image = image.(string)
	}

	// send updated user to database
	user, err = h.UserRepository.UpdateUser(user)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// reload data
	user, err = h.UserRepository.FindUserByID(user.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertUserResponse(user),
	}
	c.JSON(http.StatusOK, response)
}

func (h *handlerUser) UpdateProfile(c *gin.Context) {
	var request dto.UpdateUserRequest

	// get user data from request body
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

	// get user data from database
	user, err := h.UserRepository.FindUserByID(id)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// update fullname
	if request.FullName != "" {
		user.FullName = request.FullName
	}

	// update email
	if request.Email != "" {
		user.Email = request.Email
		user.IsEmailVerified = false
	}

	// update phone
	if request.Phone != "" {
		user.Phone = request.Phone
		user.IsPhoneVerified = false
	}

	// update address
	if request.Address != "" {
		user.Address = request.Address
	}

	// update role
	// if request.RoleID != 0 {
	// 	user.RoleID = request.RoleID
	// }

	// update image
	image, ok := c.Get("image")
	if ok {
		if user.Image != "" {
			if !helpers.DeleteFile(user.Image) {
				fmt.Println(err.Error())
			}
		}

		user.Image = image.(string)
	}

	user, err = h.UserRepository.UpdateUser(user)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// reload data
	user, err = h.UserRepository.FindUserByID(user.ID)
	if err != nil {
		response := dto.Result{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, response)
		return
	}

	// send response
	response := dto.Result{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    convertUserResponse(user),
	}
	c.JSON(http.StatusOK, response)
}

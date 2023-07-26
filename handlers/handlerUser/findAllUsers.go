package handlerUser

import (
	"go-restapi-boilerplate/dto"
	"go-restapi-boilerplate/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlerUser) FindAllUsers(c *gin.Context) {
	var (
		users       *[]models.MstUser
		err         error
		totalUser   int64
		filterQuery dto.UserFilter
	)

	// get filter data
	roleId, _ := strconv.Atoi(c.Query("roleId"))
	filterQuery.RoleID = uint(roleId)

	// get search query
	searchQuery := c.Query("search")

	// with pagination
	if c.Query("page") != "" {
		var (
			limit  int
			offset int
		)

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		// set limit (if not exist, use default limit -> 5)
		if c.Query("limit") != "" {
			limit, err = strconv.Atoi(c.Query("limit"))
			if err != nil {
				response := dto.Result{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
				}
				c.JSON(http.StatusBadRequest, response)
				return
			}
		} else {
			limit = 5

		}

		// set offset
		if page == 1 {
			offset = -1
		} else {
			offset = (page * limit) - limit
		}

		// get all users
		users, totalUser, err = h.UserRepository.FindAllUsers(limit, offset, filterQuery, searchQuery)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			}
			c.JSON(http.StatusNotFound, response)
			return
		}

		// setup and send response
		response := dto.Result{
			Status:      http.StatusOK,
			Message:     "OK",
			TotalData:   totalUser,
			TotalPages:  int(math.Ceil(float64(float64(totalUser) / float64(limit)))),
			CurrentPage: page,
			Data:        convertMultipleUserResponse(users),
		}
		c.JSON(http.StatusOK, response)
	} else { // without pagination

		// get all users
		users, totalUser, err = h.UserRepository.FindAllUsers(-1, -1, filterQuery, searchQuery)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			}
			c.JSON(http.StatusNotFound, response)
			return
		}

		// response
		response := dto.Result{
			Status:      http.StatusOK,
			Message:     "OK",
			TotalData:   totalUser,
			TotalPages:  1,
			CurrentPage: 1,
			Data:        convertMultipleUserResponse(users),
		}
		c.JSON(http.StatusOK, response)
	}
}

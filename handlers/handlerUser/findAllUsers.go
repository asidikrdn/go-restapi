package handlerUser

import (
	"go-restapi-boilerplate/db/models"
	"go-restapi-boilerplate/dto"
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
		limit, err := getLimitParam(c)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		// set offset
		offset := getOffset(page, limit)

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

func getLimitParam(c *gin.Context) (int, error) {
	if c.Query("limit") != "" {
		return strconv.Atoi(c.Query("limit"))
	}
	return 5, nil
}

func getOffset(page, limit int) int {
	if page == 1 {
		return -1
	}
	return (page * limit) - limit
}

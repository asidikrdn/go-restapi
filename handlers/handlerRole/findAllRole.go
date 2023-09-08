package handlerRole

import (
	"go-restapi/db/models"
	"go-restapi/dto"
	"go-restapi/pkg/helpers"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handlerRole) FindAllRole(c *gin.Context) {
	var (
		roles     *[]models.MstRole
		err       error
		totalRole int64
	)

	// get search query
	searchQuery := c.Query("search")

	// with pagination
	if c.Query("page") != "" {

		// get page position
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
		limit, err := helpers.GetLimitParam(c)
		if err != nil {
			response := dto.Result{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		// set offset
		offset := helpers.GetOffset(page, limit)

		// get role data from database
		roles, totalRole, err = h.RoleRepository.FindAllRole(limit, offset, searchQuery)
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
			Status:      http.StatusOK,
			Message:     "OK",
			TotalData:   totalRole,
			TotalPages:  int(math.Ceil(float64(float64(totalRole) / float64(limit)))),
			CurrentPage: page,
			Data:        convertMultipleRoleResponse(roles),
		}
		c.JSON(http.StatusOK, response)
	} else { // without pagination

		// get role data from database
		roles, totalRole, err = h.RoleRepository.FindAllRole(-1, -1, searchQuery)
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
			Status:      http.StatusOK,
			Message:     "OK",
			TotalData:   totalRole,
			TotalPages:  1,
			CurrentPage: 1,
			Data:        convertMultipleRoleResponse(roles),
		}
		c.JSON(http.StatusOK, response)
	}
}

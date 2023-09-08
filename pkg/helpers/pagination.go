package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLimitParam(c *gin.Context) (int, error) {
	if c.Query("limit") != "" {
		return strconv.Atoi(c.Query("limit"))
	}
	return 5, nil
}

func GetOffset(page, limit int) int {
	if page == 1 {
		return -1
	}
	return (page * limit) - limit
}

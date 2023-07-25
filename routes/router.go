package routes

import "github.com/gin-gonic/gin"

func RouterInit(r *gin.RouterGroup) {
	User(r)
}

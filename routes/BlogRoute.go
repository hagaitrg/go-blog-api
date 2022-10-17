package routes

import (
	"gin.com/gin/controllers"
	"github.com/gin-gonic/gin"
)

func BlogRoute(router *gin.Engine){
	router.POST("/api/v1/blogs", controllers.Create)
}
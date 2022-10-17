package routes

import (
	"gin.com/gin/controllers"
	"github.com/gin-gonic/gin"
)

func BlogRoute(router *gin.Engine){
	router.POST("/api/v1/blogs", controllers.Create)
	router.GET("/api/v1/blogs/:id", controllers.Show)
	router.PUT("api/v1/blogs/:id", controllers.Update)
}
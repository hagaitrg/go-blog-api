package main

import (
	"gin.com/gin/models"
	"gin.com/gin/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectionDatabase()
	r.GET("api/v1/blogs", controllers.Index)
	r.GET("api/v1/blogs/:id", controllers.Show)
	r.POST("api/v1/blogs", controllers.Create)

	r.Run()
}
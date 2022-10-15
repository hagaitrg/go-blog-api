package controllers

import (
	"gin.com/gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Slug string `json:"slug"`
}


// get all blogs
func Index(c *gin.Context){
	var blogs []models.Blog

	models.DB.Find(&blogs)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code": 200,
		"message": "Successfully get all blogs",
		"data" : blogs,
	})
}
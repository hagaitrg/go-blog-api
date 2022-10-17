package controllers

import (
	"gin.com/gin/models"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

// get detail blog
func Show(c *gin.Context){
	var blog models.Blog
	id := c.Param("id")

	if err:= models.DB.First(&blog, id).Error; err != nil {
		switch err{
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"success": false,
				"code": 404,
				"message": "Blog not found!",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"code": 500,
				"message": err.Error(),
			})
			return
		}
	}

	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code": 200,
		"message": "Successfully get detail blog",
		"data" : blog,
	})
}

// create blog
func Create(c *gin.Context){
	var blog models.Blog 

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"code": 400,
			"message": err.Error(),
		})
		return 
	}

	fmt.Println(blog)

	models.DB.Create(&blog)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code": 200,
		"message": "Successfully create blog",
		"data" : blog,
	})
}
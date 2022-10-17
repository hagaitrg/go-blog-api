package controllers 

import (
	"gin.com/gin/models"
	"gin.com/gin/configs"
	"gin.com/gin/handlers"
	"net/http"
	"time"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var blogCollection *mongo.Collection = configs.GetCollection(configs.DB, "go-blogs")
var validate = validator.New()

func Create(c * gin.Context){
	ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var blog models.Blog 
	defer cancel()

	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, handlers.BlogResponse{
			Success:false,
			Code: 400,
			Message: "Failed to create blog!",
			Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if validationErr := validate.Struct(&blog); validationErr != nil {
		c.JSON(http.StatusBadRequest, handlers.BlogResponse{
			Success:false,
			Code: 400,
			Message: "Failed to create blog!",
			Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	newBlog := models.Blog{
		Id: primitive.NewObjectID(),
		Title: blog.Title,
		Content: blog.Content,
		Slug: blog.Slug,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}	

	result, err := blogCollection.InsertOne(ctx,newBlog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, handlers.BlogResponse{
			Success:false,
			Code: 500,
			Message: "Failed to create blog!",
			Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, handlers.BlogResponse{
		Success:true,
		Code: 200,
		Message: "Successfully create blog!",
		Data: map[string]interface{}{"data":result}})


}
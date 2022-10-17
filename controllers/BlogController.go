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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
		Code: 201,
		Message: "Successfully create blog!",
		Data: map[string]interface{}{"data":result}})
}

func Show(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var blog models.Blog 
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	err := blogCollection.FindOne(ctx, bson.M{"_id":objId}).Decode(&blog)
	if err != nil {
		c.JSON(http.StatusNotFound, handlers.BlogResponse{
			Success:false,
			Code: 404,
			Message: "Blog Not Found!",
			Data: map[string]interface{}{"data":err.Error()}})
		return
	}

	c.JSON(http.StatusOK, handlers.BlogResponse{
		Success:true,
		Code: 200,
		Message: "Successfully get all blog!",
		Data: map[string]interface{}{"data":blog}})
}

func Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var blog models.Blog 
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(id)

	if err := c.BindJSON(&blog); err!= nil {
		c.JSON(http.StatusBadRequest, handlers.BlogResponse{
			Success:false,
			Code: 400,
			Message: "Failed!",
			Data: map[string]interface{}{"data":err.Error()}})
		return 
	}

	if validationErr := validate.Struct(&blog); validationErr != nil {
		c.JSON(http.StatusBadRequest, handlers.BlogResponse{
			Success:false,
			Code: 400,
			Message: "Validation error",
			Data: map[string]interface{}{"data":validationErr.Error()}})
		return 
	}

	update := bson.M{
		"title":blog.Title,
		"content":blog.Content,
		"slug":blog.Slug,
		"updated_at" : time.Now().UTC()}
	result, err := blogCollection.UpdateOne(ctx, bson.M{"_id":objId}, bson.M{"$set":update})
	if err != nil {
		c.JSON(http.StatusInternalServerError, handlers.BlogResponse{
			Success:false,
			Code: 500,
			Message: "Failed to update blog!",
			Data: map[string]interface{}{"data":err.Error()}})
		return 
	}

	var updateBlog models.Blog
	if result.MatchedCount == 1{
		err := blogCollection.FindOne(ctx, bson.M{"_id":objId}).Decode(&updateBlog)
		if err != nil {
			c.JSON(http.StatusInternalServerError, handlers.BlogResponse{
				Success:false,
				Code: 500,
				Message: "Failed to get blog!",
				Data: map[string]interface{}{"data":err.Error()}})
			return 
		}
	}

	c.JSON(http.StatusOK, handlers.BlogResponse{
		Success:true,
		Code: 200,
		Message: "Successfully to get blog!",
		Data: map[string]interface{}{"data":updateBlog}})
	return 
}


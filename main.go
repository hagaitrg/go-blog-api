package main

import (
	"gin.com/gin/configs"
	"gin.com/gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectionDB()

	routes.BlogRoute(router)

	router.Run("localhost:3000")
}
package main

import (
	"net/http"
	"scraper-api/db"
	"scraper-api/middlewares"
	"scraper-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	//Global Error Handler
	server.Use(middlewares.ErrorHandler())

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "route not found",
		})
	})

	routes.RegisterProductRoutes(server)

	server.Run(":8000")
}

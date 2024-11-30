package main

import (
	"net/http"
	customerrors "scraper-api/customErrors"
	"scraper-api/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//Global Error Handler
	server.Use(middlewares.ErrorHandler())

	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "route not found",
		})
	})

	server.GET("/", func(ctx *gin.Context) {
		err := customerrors.NewCustomError("dadadsads", http.StatusConflict)
		ctx.Error(err)
	})

	server.Run(":8000")
}

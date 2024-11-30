package routes

import (
	"scraper-api/services"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(server *gin.Engine) {
	server.GET("/products", services.GetAllProducts)
}

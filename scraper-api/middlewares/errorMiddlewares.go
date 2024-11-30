package middlewares

import (
	"net/http"
	customerrors "scraper-api/customErrors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		for _, err := range ctx.Errors {
			switch e := err.Err.(type) {
			case customerrors.CustomError:
				ctx.AbortWithStatusJSON(e.StatusCode, gin.H{
					"message": e.Message,
				})
			default:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"messagee": http.StatusText(http.StatusInternalServerError),
				})
			}
		}
	}
}

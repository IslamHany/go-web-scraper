package services

import (
	"net/http"
	customerrors "scraper-api/customErrors"
	"scraper-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(ctx *gin.Context) {
	limitStr, limitExists := ctx.GetQuery("limit")
	pageStr, pageExists := ctx.GetQuery("page")
	var limit, page int64
	var err error

	if limitExists {
		limit, err = strconv.ParseInt(limitStr, 10, 64)

		if err != nil {
			ctx.Error(customerrors.NewCustomError("Failed to parse limit param", http.StatusInternalServerError))
			return
		}
	}

	if pageExists {
		page, err = strconv.ParseInt(pageStr, 10, 64)

		if err != nil {
			ctx.Error(customerrors.NewCustomError("Failed to parse page param", http.StatusInternalServerError))
			return
		}
	}

	products, err := models.GetAllProducts(limit, page)

	if err != nil {
		ctx.Error(customerrors.NewCustomError("Failed to fetch events", http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

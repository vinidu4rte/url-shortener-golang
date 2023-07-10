package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vinidu4rte/url-shortener-golang/handler/response"
	"github.com/vinidu4rte/url-shortener-golang/schemas"
	"gorm.io/gorm"
)

func Redirect(ctx *gin.Context) {
	shortUrl := ctx.Params.ByName("shortUrl")

	existentRedirect := schemas.Redirect{}
	dbGetResult := db.First(&existentRedirect, "short_url = ?", shortUrl)

	if dbGetResult.Error == nil {
		ctx.Redirect(301, existentRedirect.OriginalUrl)
		return
	}

	if dbGetResult.Error == gorm.ErrRecordNotFound {
		response.HttpError(ctx, "short url dont has any redirect")
	} else {
		response.HttpError(ctx, "some error occurred in redirecting")
	}
	return
}

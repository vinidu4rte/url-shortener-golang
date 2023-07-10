package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vinidu4rte/url-shortener-golang/handler/request"
	"github.com/vinidu4rte/url-shortener-golang/handler/response"
	"github.com/vinidu4rte/url-shortener-golang/helpers"
	"github.com/vinidu4rte/url-shortener-golang/schemas"
	"gorm.io/gorm"
)

func CreateRedirect(ctx *gin.Context) {
	req := request.CreateRedirectRequest{}
	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		response.HttpError(ctx, "Malformed request.")
		return
	}

	err = req.Validate()
	if err != nil {
		response.HttpError(ctx, err.Error())
		return
	}

	shortUrl := ""
	count := 0
	existentRedirect := schemas.Redirect{}

	for true {
		shortUrl = helpers.GenerateRandomString(6)

		result := db.First(&existentRedirect, "short_url = ?", shortUrl)

		if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
			break
		}

		if count > 10 {
			response.HttpError(
				ctx,
				"An error occurred, please try again in some minutes.",
			)
			return
		}

		count += 1
	}

	redirect := schemas.Redirect{
		OriginalUrl: req.OriginalUrl,
		ShortUrl:    shortUrl,
	}

	dbCreateResult := db.Create(&redirect)

	if dbCreateResult.Error != nil {
		response.HttpError(ctx, dbCreateResult.Error.Error())
		return
	}

	response.HttpSuccess(ctx, "short url successful created", redirect)
}

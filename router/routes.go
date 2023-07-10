package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vinidu4rte/url-shortener-golang/handler"
)

func InitializeRoutes(app *gin.Engine) {
	handler.InitializeHandler()

	v1 := app.Group("/api/v1")
	{
		v1.POST("/", handler.CreateRedirect)
	}
}

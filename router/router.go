package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vinidu4rte/url-shortener-golang/handler"
)

func Handle(app *gin.Engine) {
	v1 := app.Group("/api/v1")
	{
		v1.GET("/ping", handler.PingHandler)
	}
}

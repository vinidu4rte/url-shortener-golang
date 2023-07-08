package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handle(app *gin.Engine) {
	v1 := app.Group("/api/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
}

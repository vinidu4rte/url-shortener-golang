package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vinidu4rte/url-shortener-golang/handler"
)

func Handle() {
	app := gin.Default()

	v1 := app.Group("/api/v1")
	{
		v1.GET("/ping", handler.PingHandler)
	}

	err := app.Run()

	if err != nil {
		fmt.Println("An error occurred in gin run.")
		panic(err)
	}
}

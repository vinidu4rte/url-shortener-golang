package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vinidu4rte/url-shortener-golang/router"
)

func main() {
	app := gin.Default()

	router.Handle(app)

	err := app.Run()

	if err != nil {
		fmt.Println("An error occurred in gin run.")
	}
}

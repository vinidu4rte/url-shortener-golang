package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vinidu4rte/url-shortener-golang/config"
	"github.com/vinidu4rte/url-shortener-golang/router"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	app := gin.Default()
	router.Handle(app)

	err = app.Run()

	if err != nil {
		fmt.Println("An error occurred in gin run.")
	}
}

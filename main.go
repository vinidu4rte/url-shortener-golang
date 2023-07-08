package main

import (
	"github.com/vinidu4rte/url-shortener-golang/config"
	"github.com/vinidu4rte/url-shortener-golang/router"
)

func main() {
	err := config.Init()

	if err != nil {
		panic(err)
	}

	router.Handle()
}

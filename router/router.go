package router

import (
	"github.com/gin-gonic/gin"
)

func Handle() {
	app := gin.Default()

	InitializeRoutes(app)

	app.Run()
}

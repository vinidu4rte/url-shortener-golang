package response

import "github.com/gin-gonic/gin"

func HttpError(ctx *gin.Context, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(400, gin.H{
		"message": message,
	})
}

func HttpSuccess(ctx *gin.Context, message string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, gin.H{
		"message": message,
		"data":    data,
	})
}

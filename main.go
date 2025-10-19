package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/hello", func(context *gin.Context) {
		context.JSON(200, map[string]any{
			"message": "Hello, world!",
		})
	})
	router.Run()
}

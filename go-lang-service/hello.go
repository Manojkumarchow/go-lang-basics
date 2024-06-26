package main

import "github.com/gin-gonic/gin"

func main () {
	r := gin.Default()

	r.GET("/hello", func(req *gin.Context) {

		req.JSON(200, gin.H {
			"message": "Hello world",
		})
	})

	r.Run()
}
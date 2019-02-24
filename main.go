package main

import (
	"gestbook/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/", controller.IndexGET)
	router.POST("/post", controller.ReceivePost)
	router.Run(":8080")
}

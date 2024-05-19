package router

import (
	"v-chat-back-go/handlers"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", handlers.HomeHandler)

	return r
}
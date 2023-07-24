package routes

import (
	"golang_line/app/http/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {
	router.POST("/", controllers.LineController().LineCallBack())
}

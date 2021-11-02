package router

import (
	"gin-demo/controller"
	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine) {
	r.GET("/hello", controller.TestController.Hello)
	r.GET("/error", controller.TestController.Error)
	r.GET("/tool/list", controller.Tool.List)
}

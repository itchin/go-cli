package controller

import (
	"gin-demo/service"
	"github.com/gin-gonic/gin"
)

type tool struct{}

var Tool tool

func (*tool) List(c *gin.Context) {
	list, err := service.Tool.List()
	if err != nil {
		Response.Error(c, err.Error(), 60001)
		return
	}
	Response.AddData("list", list)
	Response.Output(c)
}

package controller

import (
    "github.com/gin-gonic/gin"
)

var TestController testController

type testController struct{}

func (*testController) Hello(c *gin.Context) {
    Response.SetData("Hello World")
    Response.Output(c)
}

func (*testController) User(c *gin.Context) {
    Response.AddData("name", "itchin")
    Response.AddData("age", 28)
    Response.Output(c)
}

func (*testController) Error(c *gin.Context) {
    Response.Error(c, "this is error msg...", 60001)
}

package helper

import (
	"gin-demo/model"
	"github.com/gin-gonic/gin"
)

var responseData interface{}

var responseDataMap map[string]interface{}

const DEFAULT_ERROR_CODE = 60000

type Response struct{}

func (*Response) SetData(data interface{}) {
	responseData = data
}

func (*Response) AddData(key string, data interface{}) {
	if responseDataMap == nil {
		responseDataMap = make(map[string]interface{})
	}
	responseDataMap[key] = data
	responseData = responseDataMap
}

func (*Response) Error(c *gin.Context, msg string, code ...int) {
	err_code := DEFAULT_ERROR_CODE
	if len(code) != 0 {
		err_code = code[0]
	}
	c.JSON(200, gin.H{
		"code": err_code,
		"msg":  msg,
		"data": nil,
	})
}

func (*Response) Output(c *gin.Context) {
	response := model.NewSuccessResponse()
	c.JSON(200, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
		"data": responseData,
	})
}

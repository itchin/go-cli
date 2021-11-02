package main

import (
	"gin-demo/config"
	"gin-demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Run(r)
	r.Run(config.HTTP_PORT_STRING)
}

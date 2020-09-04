package main

import (
	"github.com/itchin/gin-demo/helper/config"
	"github.com/itchin/gin-demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Run(r)
	r.Run(config.HTTP_PORT_STRING)
}

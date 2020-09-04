package orm

import (
	"fmt"
	"github.com/itchin/gin-demo/helper/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var x *xorm.Engine

func init() {
	var err error
	port := ""
	if config.DB_PORT != 0 {
		port = fmt.Sprintf(":%d", config.DB_PORT)
	}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=utf8", config.DB_USER, config.DB_PWD, config.DB_HOST, port, config.DB_NAME)
	x, err = xorm.NewEngine(config.DB_TYPE, dataSourceName)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
}

func Engine() *xorm.Engine {
	return x
}

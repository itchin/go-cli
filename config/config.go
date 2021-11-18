package config

import (
	"gopkg.in/ini.v1"
	"fmt"
	"os"
	"strings"
)

var (
	// DEMO_PATH 模板框架所在目录
	DEMO_PATH string = "D:\\project\\go-cli\\gin-demo"
	// DEMO_PKG_NAME 框架原包名
	DEMO_PKG_NAME string = "github.com/itchin/gin-demo"
	// PKG_NAME 新建项目的包名
	PKG_NAME string = "myapp"
	// PATH 新建项目的路径，如果为空，则默认在go-cli目录下
	PATH string = ""
)

/**
 * 读取config.ini的环境配置
 */
func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read config.ini: %v", err)
		os.Exit(1)
	}

	appConfigInit(cfg.Section("APP"))
}

/**
 * 读取配置文件中APP项的配置
 */
func appConfigInit(section *ini.Section) {
	demoPath := section.Key("DEMO_PATH").String()
	if demoPath != "" {
		DEMO_PATH = demoPath
	}

	demoPkgName := section.Key("DEMO_PKG_NAME").String()
	if demoPkgName != "" {
		DEMO_PKG_NAME = demoPkgName
	}

	pkgName := section.Key("PKG_NAME").String()
	if pkgName != "" {
		PKG_NAME = pkgName
	}

	path := section.Key("PATH").String()
	if path != "" {
		//windows路径将反斜杠转为斜杠
		path = strings.ReplaceAll(path, "\\", "/")
		//如果不是斜杠结尾，加上它
		if path[len(path)-1:] != "/" {
			path += "/"
		}
		PATH = path
	}
}

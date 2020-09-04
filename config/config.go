package config

import (
    "fmt"
    "gopkg.in/ini.v1"
    "os"
)

var (
    DEMO_PATH string
    DEMO_PKG_NAME string
    PKG_NAME string
)

func init()  {
    cfg, err := ini.Load("env.ini")
    if err != nil {
        fmt.Printf("Fail to read config.ini: %v", err)
        os.Exit(1)
    }

    section := cfg.Section("")

    DEMO_PATH = section.Key("DEMO_PATH").String()
    DEMO_PKG_NAME = section.Key("DEMO_PKG_NAME").String()
    PKG_NAME = section.Key("PKG_NAME").String()
}

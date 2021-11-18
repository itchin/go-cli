package config

import (
    "fmt"
    "os"
    "strconv"

    "gopkg.in/ini.v1"
)

var (
    HTTP_PORT int
    HTTP_PORT_STRING string

    DB_TYPE string
    DB_HOST string
    DB_PORT int
    DB_NAME string
    DB_USER string
    DB_PWD string
)

func init() {
    cfg, err := ini.Load("config.ini")
    if err != nil {
        fmt.Printf("Fail to read config.ini: %v", err)
        os.Exit(1)
    }

    appConfigInit(cfg.Section("APP"))
    dbConfigInit(cfg.Section("DB"))
}

func appConfigInit(section *ini.Section) {
    port, err := section.Key("HTTP_PORT").Int()
    if err == nil {
        HTTP_PORT = port
        HTTP_PORT_STRING = ":" + strconv.Itoa(port)
    }
}

func dbConfigInit(section *ini.Section) {
    DB_TYPE = section.Key("TYPE").String()
    DB_HOST = section.Key("HOST").String()
    port, err := section.Key("PORT").Int()
    if err == nil {
        DB_PORT = port
    }
    DB_NAME = section.Key("NAME").String()
    DB_USER = section.Key("USER").String()
    DB_PWD = section.Key("PWD").String()
}

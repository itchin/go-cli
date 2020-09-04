package main

import (
    "github.com/itchin/go-cli/config"
    "github.com/itchin/go-cli/utils"
    "fmt"
    "github.com/urfave/cli"
    "os"
)


func main() {
    app := cli.NewApp()
    app.Name = "go-cli" // 指定程序名称
    app.Usage = "golang application installer" //  程序功能描述
    app.UsageText = "go-cli api --dir [your-project-dir] --package [package name]"
    app.Description = "usage example: go-cli api --dir dirName --package github.com/itchin/go-cli"
    app.HideHelp = true
    app.Commands = []cli.Command{
        {
            Name:     "api",         //命令名称
            Usage:    "download frame project...", // 命令描述
            Category: "arithmetic",  // 命令所属的类别
            Action: apiAction,
            //Action: func(c *cli.Context) error { // 函数入口
            //    fmt.Println("this app name is:", c.String("appname"))
            //    return nil
            //},
            Flags: []cli.Flag{
                &cli.StringFlag{
                    Name:    "dir",
                    Value:   "app-demo",
                    Usage:   "directory name",
                },
                &cli.StringFlag{
                    Name: "package",
                    Value: "app-pkg",
                    Usage: "package name",
                },
            },
        },
    }

    app.Run(os.Args)
}

func apiAction(c *cli.Context) error {
    dir := c.String("dir")
    if utils.IsDirExists(dir) == false {
        err := os.MkdirAll(dir, os.ModePerm)
        if err != nil {
            panic("创建路径失败：" + err.Error())
        }
    }
    err := utils.CopyDir(config.DEMO_PATH, c.String("dir"), config.DEMO_PKG_NAME, config.PKG_NAME)
    if err == nil {
        fmt.Println("app install success...")
    } else {
        panic(err)
    }
    return nil
}

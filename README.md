基于 https://github.com/urfave/cli 进行开发，快速构建app项目脚手架。

### 二进制下载

https://github.com/itchin/go-cli/releases

### 部署说明

需安装golang开发环境

```
git clone https://github.com/itchin/go-cli

# 下载所需组件
go mod vendor

# 配置env.ini后，安装go-cli
go install go-cli.go

# 构建项目脚手架
go-cli api --dir myapp --package github.com/itchin/myapp
```

### go-cli 参数说明
--dir: 应用所在的目录名
--package: 应用go modules命名空间

### env.ini 配置

复制 env.exam.ini 为 env.ini

```
# 模板框架所在目录，程序中过滤目录中的.git目录
DEMO_PATH  = "D:/goProject/go-cli/gin-demo"
# 框架原包名
DEMO_PKG_NAME = "github.com/itchin/gin-demo"
# 进行替换的包名
PKG_NAME = "user-service"
```
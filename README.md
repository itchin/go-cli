根据模板框架，快速生成app项目。

### 二进制下载

https://github.com/itchin/go-cli/releases

### 部署说明

需安装golang开发环境

```
git clone https://github.com/itchin/go-cli

# 下载所需组件
go mod vendor

# 修改config/config.go中配置，安装go-cli
go install go-cli.go

# 生成项目
go-cli api --dir myapp --package github.com/itchin/myapp
```

### go-cli 参数说明
--dir: 应用所在的目录名

--package: 应用go modules命名空间

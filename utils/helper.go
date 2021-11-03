package utils

import (
    "errors"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

// IsDirExists
// 判断目录是否存在
func IsDirExists(fileAddr string)bool{
    s,err := os.Stat(fileAddr)
    if err != nil{
        return false
    }
    return s.IsDir()
}

// CopyDir
/**
 * 拷贝文件夹,同时拷贝文件夹中的文件
 * @param srcPath  		需要拷贝的文件夹路径: D:/test
 * @param destPath		拷贝到的位置: D:/backup/
 */
func CopyDir(srcPath string, destPath string, demoPkgName string, pkgName string) error {
    //检测目录正确性
    if srcInfo, err := os.Stat(srcPath); err != nil {
        //路径不存在
        fmt.Println(err.Error())
        return err
    } else if !srcInfo.IsDir() {
        //路径为文件，不是目录
        e := errors.New("scrPath is not a dir")
        fmt.Println(e.Error())
        return e
    }
    if destInfo, err := os.Stat(destPath); err != nil {
        fmt.Println(err.Error())
        return err
    } else if !destInfo.IsDir() {
        e := errors.New("destInfo not exist")
        fmt.Println(e.Error())
        return e
    }

    err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
        if f == nil {
            return err
        }
        if !f.IsDir() {
            path = strings.Replace(path, "\\", "/", -1)
            srcPath = strings.Replace(srcPath, "\\", "/", -1)
            // 过滤模板框架的.git目录
            if strings.Contains(path, "/.git/") {
                return nil
            }
            destNewPath := strings.Replace(path, srcPath, destPath, -1)
            //fmt.Println(path, srcPath, destPath)
            //fmt.Println("复制文件:" + path + " 到 " + destNewPath)
            //fmt.Println(path, destNewPath, demoPkgName, pkgName)
            err = copyAndRelaceFile(path, destNewPath, demoPkgName, pkgName)
            if err != nil {
                panic(err)
            }
        }
        return nil
    })
    if err != nil {
        fmt.Printf(err.Error())
    }
    return err
}

//生成目录并拷贝文件
func copyAndRelaceFile(src, dest, demoPkgName, pkgName string) (err error) {
    //分割path目录
    destSplitPathDirs := strings.Split(dest, "/")

    //检测时候存在目录
    destSplitPath := ""
    for index, dir := range destSplitPathDirs {
        if index < len(destSplitPathDirs)-1 {
            destSplitPath = destSplitPath + dir + "/"
            b, _ := pathExists(destSplitPath)
            if b == false {
                //fmt.Println("创建目录:" + destSplitPath)
                //创建目录
                err := os.Mkdir(destSplitPath, os.ModePerm)
                if err != nil {
                    fmt.Println(err)
                }
            }
        }
    }
    dstFile, err := os.Create(dest)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer dstFile.Close()

    bytes, err := ioutil.ReadFile(src)
    content := string(bytes)
    content = strings.ReplaceAll(content, demoPkgName, pkgName)
    dstFile.WriteString(content)

    return
}

//检测文件夹路径时候存在
func pathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}


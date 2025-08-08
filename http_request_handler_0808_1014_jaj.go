// 代码生成时间: 2025-08-08 10:14:46
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo-pop/v2/pop/popmw"
    "github.com/unrolled/secure"
    "log"
)

// App 是Buffalo应用的主要结构体
var App *buffalo.App

// DB 是数据库连接，用于操作数据库
var DB buffalo.DB

// main 函数是程序的入口点
func main() {
    // 创建一个Buffalo应用
    App = buffalo.Automatic(buffalo.Options{
       Preinitializers: []buffalo.Preinitializer{
           // 在Buffalo应用之前运行的预初始化器
           popmw.NullsAsPointers,
       },
       Middleware: []buffalo.MiddleWare{
           // 添加中间件
           secure.New(secure.Options{
               FrameDeny: true,
              ContentTypeNosniff: true,
           })},
    )
    
    // 设置数据库连接
    DB = buffalo.DB
    
    // 注册路由
    App.GET("/", HomeHandler)
    App.POST("/data", DataHandler)
    
    // 启动服务器
    if err := App.Serve(); err != nil {
        log.Fatal(err)
    }
}

// HomeHandler 是首页的HTTP请求处理器
func HomeHandler(c buffalo.Context) error {
    // 渲染首页模板
    return c.Render(200, buffalo.HTML("index.html"))
}

// DataHandler 是处理POST请求的HTTP请求处理器
func DataHandler(c buffalo.Context) error {
    // 从请求中获取数据
    data := new(DataModel)
    if err := c.Bind(data); err != nil {
        // 绑定数据出错
        return err
    }
    
    // 保存数据到数据库
    if err := DB.Create(data); err != nil {
        // 数据库操作出错
        return err
    }
    
    // 返回成功的响应
    return c.Render(201, buffalo.JSON(data))
}

// DataModel 是数据模型，用于表示数据库中的记录
type DataModel struct {
    // 定义字段
    ID    uint   `db:"id"`
    Name  string `db:"name"`
    Value string `db:"value"`
}

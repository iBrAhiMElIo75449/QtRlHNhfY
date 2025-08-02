// 代码生成时间: 2025-08-02 20:57:45
package main

import (
    "buffalo"
    "github.com/markbates/inflect"
    "log"
    "math/rand"
    "time"
)

// RandomNumberGeneratorHandler handles the request to generate a random number
func RandomNumberGeneratorHandler(c buffalo.Context) error {
    // 生成随机数的范围
    min := 1
    max := 100

    // 生成随机数
    rand.Seed(time.Now().UnixNano())
    randomNumber := rand.Intn(max-min) + min

    // 返回JSON响应
    return c.Render(200,(buffalo.R{
        "randomNumber": randomNumber,
    }))
}

// main function to setup the buffalo application
func main() {
    app := buffalo.New(inflect.Underscore("random_number_generator"))

    // 定义路由
    app.GET("/random-number", RandomNumberGeneratorHandler)

    // 启动服务器
    log.Fatal(app.Start(":3000"))
}

// 代码生成时间: 2025-09-16 08:23:30
@author: Your name
@date: 2023-04-01
*/

package main

import (
    "os"
    "log"
    "path/filepath"
    "github.com/unidoc/unioffice/excel"
    "github.com/markbates/buffalo"
)

// ExcelGenerator 结构体，用于生成Excel表格
type ExcelGenerator struct {
    // 定义需要的字段
}

// NewExcelGenerator 创建一个新的ExcelGenerator实例
func NewExcelGenerator() *ExcelGenerator {
    return &ExcelGenerator{}
}

// GenerateExcel 生成Excel表格
func (g *ExcelGenerator) GenerateExcel(rows [][]string) error {
    // 创建一个新的Excel文档
    f, err := excel.NewFile()
    if err != nil {
        return err
    }
    defer f.Close()

    // 创建一个新的工作表
    sheet, err := f.AddSheet("Sheet1")
    if err != nil {
        return err
    }

    // 添加数据到工作表
    for i, row := range rows {
        for j, value := range row {
            if err := sheet.SetCellValue(j, i, value); err != nil {
                return err
            }
        }
    }

    // 保存Excel文档
    return f.SaveAs("example.xlsx")
}

// App 结构体，用于BUFFALO框架
type App struct {
    Generator *ExcelGenerator
}

// NewApp 创建一个新的App实例
func NewApp() *App {
    return &App{
        Generator: NewExcelGenerator(),
    }
}

func main() {
    // 创建一个BUFFALO应用
    app := buffalo.Automatic(buffalo.Options{
        CustomApp: func() *buffalo.App {
            app := buffalo.New(buffalo.Options{})
            app.Middleware.Add(
                negroni.NewRecovery(),
                negroni.NewLogger(),
                negroni.NewStatic(http.FileServer(http.Dir("github.com/markbates/buffalo"))),
            )
            app.GET("/generate", NewApp().Handler)
            return app
        },
    })
    app.Serve()
}

// Handler 处理生成Excel表格的请求
func (a *App) Handler(c buffalo.Context) error {
    // 定义Excel表格的数据
    data := [][]string{
        {"Name", "Age", "City"},
        {"Alice", "30", "New York"},
        {"Bob", "25", "Los Angeles"},
        {"Charlie", "35", "Chicago"},
    }

    // 生成Excel表格
    if err := a.Generator.GenerateExcel(data); err != nil {
        log.Printf("Error generating Excel: %v", err)
        return c.Render(500, r.JSON(map[string]string{"error": "Failed to generate Excel"}))
    }

    // 返回成功响应
    return c.Render(200, r.JSON(map[string]string{"message": "Excel generated successfully"}))
}
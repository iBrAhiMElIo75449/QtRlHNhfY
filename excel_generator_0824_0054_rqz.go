// 代码生成时间: 2025-08-24 00:54:55
package main

import (
    "archive/zip"
    "encoding/xml"
    "fmt"
    "os"
    "time"

    "github.com/tealeg/xlsx/v3"
    "github.com/markbates/buffalo"
)

// ExcelGenerator 是一个用于生成Excel文件的结构体
type ExcelGenerator struct {
    // SheetName 是工作表的名称
    SheetName string
    // Data 是要写入工作表的数据
    Data [][]string
}

// NewExcelGenerator 创建一个新的ExcelGenerator实例
func NewExcelGenerator(sheetName string, data [][]string) *ExcelGenerator {
    return &ExcelGenerator{
        SheetName: sheetName,
        Data: data,
    }
}

// GenerateExcel 创建一个新的Excel文件并写入数据
func (e *ExcelGenerator) GenerateExcel() (*os.File, error) {
    // 创建一个新的Excel文件
    file, err := xlsx.NewFile()
    if err != nil {
        return nil, err
    }

    // 创建一个新的工作表
    sheet, err := file.AddSheet(e.SheetName)
    if err != nil {
        return nil, err
    }

    // 将数据写入工作表
    for _, row := range e.Data {
        if err := sheet.AddRow(row); err != nil {
            return nil, err
        }
    }

    // 将Excel文件保存到临时文件
    tempFile, err := os.CreateTemp("
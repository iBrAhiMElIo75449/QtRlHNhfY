// 代码生成时间: 2025-08-02 02:07:49
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo/generators"
    "github.com/markbates/inflect"
    "log"
)

// 定义一个结构体，用于生成自动化测试套件
type TestSuiteGenerator struct {
    *generators.Generator
}

// NewTestSuiteGenerator 创建一个新的测试套件生成器
func NewTestSuiteGenerator() *TestSuiteGenerator {
    return &TestSuiteGenerator{
        Generator: generators.Generator{},
    }
}

// Generate 执行测试套件的生成
func (g *TestSuiteGenerator) Generate(f *buffalo.File) error {
    // 检查文件是否已经存在
    if f.Exists() {
        return nil
    }

    // 生成测试套件内容
    content := `package {{module}}_test

import (
    . "github.com/onsi/ginkgo"
    "{{module}}"
    . "github.com/onsi/gomega"
)

// TestSuite 测试套件
var TestSuite *testing.T

// SetupSuite 测试套件设置
var SetupSuite func()

// TearDownSuite 测试套件清理
var TearDownSuite func()

// BeforeSuite 测试集前执行
var BeforeSuite func()

// AfterSuite 测试集后执行
var AfterSuite func()

// SpecsBeforeSuite 测试集前执行
var SpecsBeforeSuite []func()

// SpecsAfterSuite 测试集后执行
var SpecsAfterSuite []func()

var _ = SynchronizedBeforeSuite(SetupSuite, TearDownSuite)
var _ = SynchronizedAfterSuite(BeforeEach, AfterEach)

// TestTestSuite 测试套件测试
func TestTestSuite(t *testing.T) {
    TestSuite = t
    RegisterFailHandler(Fail)
    RunSpecs(t, "{{module}} Suite")
}
`

    // 替换模板变量
    content = f.Replace(content, map[string]string{
        "{{module}}": "your_module_name",
    })

    // 写入文件
    if err := f.Write(content); err != nil {
        return err
    }

    return nil
}

// 运行程序
func main() {
    generator := NewTestSuiteGenerator()
    file := buffalo.NewFile("test_suite_test.go")
    if err := generator.Generate(file); err != nil {
        log.Fatal(err)
    }
}

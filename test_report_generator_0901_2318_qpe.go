// 代码生成时间: 2025-09-01 23:18:29
package main

import (
    "buffalo"
    "github.com/markbates/pkger"
    "os"
    "strings"
    "text/template"
)

// TestReport represents the data to be rendered in the template.
type TestReport struct {
    Tests      []TestResult
    Timestamp string
}

// TestResult represents an individual test result.
type TestResult struct {
    Name      string
    Duration  string
    Status    string
    Message   string
}

// NewTestReport creates a new TestReport with the current timestamp.
func NewTestReport() *TestReport {
    return &TestReport{
        Timestamp: CurrentTimestamp(),
    }
}

// CurrentTimestamp returns the current timestamp in a human-readable format.
func CurrentTimestamp() string {
    return buffalo.Now().Format("2006-01-02 15:04:05")
}

// AddTestResult adds a test result to the report.
func (tr *TestReport) AddTestResult(name string, duration string, status string, message string) {
    tr.Tests = append(tr.Tests, TestResult{
        Name:      name,
        Duration:  duration,
        Status:    status,
        Message:   message,
    })
}

// GenerateReport generates the test report HTML.
func GenerateReport(tr *TestReport) (string, error) {
    tmpl, err := template.New("report").Funcs(template.FuncMap{
        "formatDuration": func(d string) string {
            return strings.ToUpper(d)
        },
    }).ParseFiles("templates/report.html")
    if err != nil {
        return "", err
    }
    buf := new(strings.Builder)
    if err := tmpl.ExecuteTemplate(buf, "report", tr); err != nil {
        return "", err
    }
    return buf.String(), nil
}

func main() {
    app := buffalo.Buffalo(buffalo.Options{})
    app.GET("/report", func(c buffalo.Context) error {
        tr := NewTestReport()
        // Simulating test results
        tr.AddTestResult("Test1", "100ms", "Passed", "Test 1 passed successfully")
        tr.AddTestResult("Test2", "200ms", "Failed", "Test 2 failed with error")

        report, err := GenerateReport(tr)
        if err != nil {
            return c.Error(500, err)
        }

        return c.Render(200, buffalo.HTML(render.Options{"root"}).String(report))
    })

    // Start the application
    app.Serve()
}
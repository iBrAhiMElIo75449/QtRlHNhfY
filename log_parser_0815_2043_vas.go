// 代码生成时间: 2025-08-15 20:43:27
@author Your Name
*/

package main

import (
    "buffalo"
    "buffalo/buffalo/plugins"
    "buffalo/render"
    "github.com/markbates/pkg/log"
    "os"
    "path/filepath"
    "strings"
)

// LogParser is the main struct for the log parser tool.
type LogParser struct {
    // Path to the log file.
    FilePath string
    // Extensions to parse.
    Extensions []string
}

// NewLogParser initializes a new LogParser instance.
func NewLogParser(filePath string, extensions []string) *LogParser {
    return &LogParser{
        FilePath:  filePath,
        Extensions: extensions,
    }
}

// Parse parses the log file and performs necessary operations.
func (l *LogParser) Parse() ([]string, error) {
    // Check if the file exists.
    if _, err := os.Stat(l.FilePath); os.IsNotExist(err) {
        return nil, err
    }

    // Read the file content.
    fileContent, err := os.ReadFile(l.FilePath)
    if err != nil {
        return nil, err
    }

    // Split the content into lines.
    lines := strings.Split(string(fileContent), "
")

    // Filter lines based on the provided extensions.
    var filteredLines []string
    for _, line := range lines {
        for _, extension := range l.Extensions {
            if strings.Contains(line, extension) {
                filteredLines = append(filteredLines, line)
            break
            }
        }
    }

    return filteredLines, nil
}

// main function to set up BUFFALO and run the application.
func main() {
    // Set up the BUFFALO application.
    app := buffalo.New(buffalo.Options{
        PreWares: []buffalo.PreWare{
            logMiddleware, // Add log middleware for logging.
        },
    })

    // Define the route for parsing logs.
    app.GET("/parse", func(c buffalo.Context) error {
        // Extract file path and extensions from query parameters.
        filePath := c.Param("file")
        extensions := c.Request().URL.Query().Get("extensions")

        // Create a new LogParser instance.
        parser := NewLogParser(filePath, strings.Split(extensions, ","))

        // Parse the log file and handle errors.
        lines, err := parser.Parse()
        if err != nil {
            log.Error(err)
            return render.String(c, "error", "Log parsing failed.")
        }

        // Return the parsed lines as a JSON response.
        return render.JSON(c, lines)
    })

    // Run the BUFFALO application.
    app.Serve()
}

// logMiddleware is a middleware function for logging requests.
func logMiddleware(next buffalo.Handler) buffalo.Handler {
    return func(c buffalo.Context) error {
        // Log the request.
        log.Info(c.Request().Method, " ", c.Request().URL.Path)
        return next(c)
    }
}
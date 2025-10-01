// 代码生成时间: 2025-10-01 21:31:09
package main

import (
    "os"
    "log"
    "time"
    "bufio"
    "context"
    "github.com/gobuffalo/buffalo"
)

// ErrorLogger is a struct that holds the necessary info for logging errors
type ErrorLogger struct {
    context.Context
    Logger *log.Logger
}

// NewErrorLogger initializes a new ErrorLogger with a standard logger
// that logs to the standard logger file
func NewErrorLogger(ctx context.Context) *ErrorLogger {
    file, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open error log file: %v", err)
    }
    logger := log.New(file, "ERROR: ", log.LstdFlags)
    return &ErrorLogger{
        Context: ctx,
        Logger:  logger,
    }
}

// LogError logs an error message with a timestamp
func (el *ErrorLogger) LogError(err error) {
    if err != nil {
        el.Logger.Printf("%s - %s
", time.Now().Format(time.RFC3339), err.Error())
    }
}

// ErrorResponseWriter is a middleware that logs errors from requests
func ErrorResponseWriter(c buffalo.Context) error {
    // Get the error logger from the context
    el, ok := c.Value("errorLogger").(*ErrorLogger)
    if !ok {
        // If not found, create a new one
        el = NewErrorLogger(c)
        c.Set("errorLogger", el)
    }
    // Proceed to the next middleware
    err := c.Next()
    // Log the error if present
    if err != nil {
        el.LogError(err)
    }
    return err
}

// main function to setup Buffalo and start the server
func main() {
    app := buffalo.Automatic()

    // Use the ErrorResponseWriter middleware
    app.Use(ErrorResponseWriter)

    // Define a simple route to test error logging
    app.GET("/error", func(c buffalo.Context) error {
        return buffalo.NewError("Something went wrong!")
    })

    // Run the application
    app.Serve()
}

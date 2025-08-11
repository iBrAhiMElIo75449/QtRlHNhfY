// 代码生成时间: 2025-08-11 14:49:37
package main

import (
    "archive/zip"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/middleware"
    "github.com/gobuffalo/buffalo/worker"
)

// DecompressWorker handles the decompression of a file
type DecompressWorker struct {
    // Params contains the file path to be decompressed
    Params map[string]string
}

// Work decompresses the file and returns a result
func (dw *DecompressWorker) Work() (string, error) {
    file, ok := dw.Params["file"]
    if !ok || file == "" {
        return "", ErrInvalidFile
    }

    // Ensure the file is a zip file
    if !strings.HasSuffix(file, ".zip") {
        return "", ErrNotZipFile
    }

    // Create a buffer to store the decompressed files
    targetDir := file + "_decompressed"
    if _, err := os.Stat(targetDir); os.IsNotExist(err) {
        os.MkdirAll(targetDir, 0755)
    }

    // Open the zip file
    reader, err := zip.OpenReader(file)
    if err != nil {
        return "", err
    }
    defer reader.Close()

    // Iterate through the files in the zip
    for _, f := range reader.File {
        // Create the full file path
        filePath := filepath.Join(targetDir, f.Name)
        
        // Check for file existence to avoid overwriting
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            // Create the directory structure if necessary
            if f.FileInfo().IsDir() {
                os.MkdirAll(filePath, 0755)
            } else {
                // Create the file
                file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, f.Mode())
                if err != nil {
                    return "", err
                }
                
                // Extract the file content
                fileInZip, err := f.Open()
                if err != nil {
                    file.Close()
                    return "", err
                }
                defer fileInZip.Close()
                
                // Copy the content from the zip file to the new file
                _, err = io.Copy(file, fileInZip)
                if err != nil {
                    file.Close()
                    return "", err
                }
                file.Close()
            }
        } else {
            return "", ErrFileExists
        }
    }

    // Return the path to the decompressed files
    return targetDir, nil
}

// ErrInvalidFile indicates an invalid file error
var ErrInvalidFile = errors.New("invalid file path")

// ErrNotZipFile indicates a non-zip file error
var ErrNotZipFile = errors.New("not a zip file")

// ErrFileExists indicates a file already exists error
var ErrFileExists = errors.New("file already exists")

// NewDecompressWorker returns a new instance of DecompressWorker
func NewDecompressWorker(params map[string]string) worker.Worker {
    return &DecompressWorker{Params: params}
}

// App is the main application
type App struct{
    *buffalo.App
}

// Start runs the application
func (a *App) Start(address string) error {
    a.ServeFiles("/", assets.NewFileSystem())
    // Register routes
    a.GET("/decompress", a.decompressHandler)
    a.POST("/decompress", a.decompressHandler)
    return a.App.Serve(address)
}

// decompressHandler handles the decompression request
func (a *App) decompressHandler(c buffalo.Context) error {
    // Get the file from the request
    file, err := c.File("file")
    if err != nil {
        return buffalo.NewError("Invalid file")
    }
    defer file.Close()
    
    // Save the file to the server
    targetFile, err := os.Create(file.Filename)
    if err != nil {
        return err
    }
    defer targetFile.Close()
    _, err = io.Copy(targetFile, file)
    if err != nil {
        return err
    }
    
    // Decompress the file
    dw := NewDecompressWorker(map[string]string{"file": file.Filename})
    result, err := dw.Work()
    if err != nil {
        return buffalo.NewError("Decompression failed")
    }
    
    // Return the result
    return c.Render(200, r.Auto.JSON(result))
}

func main() {
    // Create the application
    app := &App{buffalo.NewApp()

    // Add middleware
    app.Use(middleware.ParameterLoggerDefault())
    app.Use(middleware BUFFALO回收站
    app.Use(middleware.PopTransactionMiddleWare())
    app.Use(middleware.Renderer(rendererOptions))

    // Start the application
    err := app.Start(":3000")
    if err != nil {
        log.Fatal(err)
    }
}

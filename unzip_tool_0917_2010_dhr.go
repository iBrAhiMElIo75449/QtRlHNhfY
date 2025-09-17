// 代码生成时间: 2025-09-17 20:10:49
 * documentation. It follows GoLang best practices for maintainability and extensibility.
 */

package main

import (
    "bufio"
    "compress/gzip"
    "flag"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gobuffalo/buffalo"
)

// Decompress takes a reader and writes the uncompressed data to a writer
func Decompress(reader io.Reader, writer io.Writer) error {
    gzipReader, err := gzip.NewReader(reader)
    if err != nil {
        return fmt.Errorf("failed to create gzip reader: %w", err)
    }
    defer gzipReader.Close()

    _, err = io.Copy(writer, gzipReader)
    return err
}

// Handler responds to the HTTP request to decompress a file
func Handler(c buffalo.Context) error {
    file, err := c.Param("file")
    if err != nil {
        return buffalo.NewError("Invalid file parameter")
    }

    // Check if the file exists
    if _, err := os.Stat(file); os.IsNotExist(err) {
        return buffalo.NewError("File does not exist")
    }

    src, err := os.Open(file)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer src.Close()

    dest := file + ".unzipped"
    dst, err := os.Create(dest)
    if err != nil {
        return fmt.Errorf("failed to create destination file: %w", err)
    }
    defer dst.Close()

    if err := Decompress(src, dst); err != nil {
        return fmt.Errorf("failed to decompress file: %w", err)
    }

    c.Response().WriteHeader(http.StatusOK)
    return nil
}

func main() {
    app := buffalo.Automatic(buffalo.Options{})
    app.GET("/unzip/{file}", Handler)
    defer app.Serve()
}
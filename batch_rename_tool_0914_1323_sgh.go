// 代码生成时间: 2025-09-14 13:23:23
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// RenameBatch defines the batch renaming operation.
type RenameBatch struct {
    // BasePath is the directory where files are located.
    BasePath string
    // Prefix is the new prefix to add to the filenames.
    Prefix string
    // Counter is the initial number to start renaming from.
    Counter int
}

// NewRenameBatch creates a new RenameBatch instance with the given parameters.
func NewRenameBatch(basePath, prefix string, counter int) *RenameBatch {
    return &RenameBatch{
        BasePath: basePath,
        Prefix: prefix,
        Counter: counter,
    }
}

// Run performs the batch renaming operation.
func (rb *RenameBatch) Run() error {
    files, err := os.ReadDir(rb.BasePath)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }
    for _, file := range files {
        if file.IsDir() {
            continue // Skip directories.
        }
        oldPath := filepath.Join(rb.BasePath, file.Name())
        newPath := filepath.Join(rb.BasePath, fmt.Sprintf("%s%03d.%s", rb.Prefix, rb.Counter, strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))))
        if err := os.Rename(oldPath, newPath); err != nil {
            return fmt.Errorf("failed to rename file %s to %s: %w", oldPath, newPath, err)
        }
        rb.Counter++ // Increment counter for the next file.
    }
    return nil
}

func main() {
    // Example usage of the batch renaming tool.
    basePath := "./files" // Replace with your directory path.
    prefix := "new_" // Replace with your preferred prefix.
    counter := 1
    batch := NewRenameBatch(basePath, prefix, counter)
    if err := batch.Run(); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Batch renaming completed successfully.")
}